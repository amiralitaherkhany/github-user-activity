package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
	fmt.Println("your username is", userName)
	numberOfActivities, err := getUserGithubActivity(userName)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Number of activities is", numberOfActivities)
}

func getUserGithubActivity(username string) (uint64, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("Response failed with status code: %d\n", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	var data []interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, errors.New("error in parsing data")
	}
	return uint64(len(data)), nil

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
