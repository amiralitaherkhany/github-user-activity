package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/amiralitaherkhany/github-user-activity/entity"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	userName, err := getUserName()
	if err != nil {
		log.Fatalln(err)
	}
	userActivities, err := getUserGithubActivities(userName)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(len(*userActivities), "Events:")
	cm := categorizeGithubEvents(userActivities)
	for eventType, eventCount := range cm {
		fmt.Printf("%s -> %d \n", eventType, eventCount)
	}
}

func categorizeGithubEvents(g *[]entity.GithubEvent) map[string]uint {
	categories := make(map[string]uint)
	for _, event := range *g {
		_, ok := categories[event.Type]
		if ok {
			categories[event.Type]++
		} else {
			categories[event.Type] = 1
		}
	}
	return categories
}

func getUserGithubActivities(username string) (*[]entity.GithubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Response failed with status code: %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error in reading from response: %s\n", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var data []entity.GithubEvent
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, errors.New("error in parsing json data")
	}

	return &data, nil
}

func getUserName() (string, error) {
	switch len(os.Args) {
	case 2:
		return os.Args[1], nil
	case 1:
		return "", errors.New("username is required")
	default:
		return "", errors.New("invalid number of flags")
	}
}
