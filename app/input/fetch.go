package input

import (
	"fmt"
	"io"
	"net/http"
)

var url = "https://adventofcode.com/%d/day/%d/input"

func FetchInput(day int, year int) ([]byte, error) {
	token, err := Token()
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
