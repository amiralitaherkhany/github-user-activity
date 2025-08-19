package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	userName, err := getUserName()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("your username is", userName)
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
