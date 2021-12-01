package app

import (
	"fmt"
	"github.com/charlesbourget/aocprep/app/fs"
	"time"
)

func Start(day int, year int, workDir string) {
	fmt.Println("Advent of Code Preparator 🎅")
	fmt.Printf("Preparing setup for day %d\n", day)

	isDateValid, err := validateDate(day, year)
	if err != nil {
		fmt.Println("Error while validating Date. Make sure the date is valid.", err)
		return
	}
	if !isDateValid {
		fmt.Println("Too early please wait for midnight EST")
		return
	}

	name := fmt.Sprintf("day%d", day)
	path := workDir
	if path[len(path)-1] != '/' {
		path += "/"
	}

	fs.CreateStructure(path, name, year, day)

	fmt.Printf("Structure for day %d created! 🚀\n", day)
}

func validateDate(day int, year int) (bool, error) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return false, err
	}
	requestedDateEST := time.Date(year, 12, day, 0, 0, 0, 0, location)

	return time.Now().In(location).After(requestedDateEST), nil
}
