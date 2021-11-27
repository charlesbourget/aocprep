package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

var url = "https://adventofcode.com/%d/day/%d/input"

func FetchInput(day int, year int) ([]byte, error) {
	token, err := loadEnv()
	if err != nil {
		return []byte{}, err
	}

	urlFormatted := fmt.Sprintf(url, year, day)
	client := http.Client{}
	req, err := http.NewRequest("GET", urlFormatted, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("cookie", token)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	var input []byte
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, err
		}

		input = bodyBytes
	}

	return input, nil
}

func loadEnv() (string, error) {
	err := godotenv.Load("resources/.env")
	if err != nil {
		return "", err
	}

	return os.Getenv("AOC_SESSION"), nil
}
