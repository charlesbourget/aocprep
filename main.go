package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/charlesbourget/aocprep/app"
	"os"
	"strconv"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: prepare <day> <work dir>")
		flag.PrintDefaults()
	}

	day, year, workDir, err := parse()
	if err != nil {
		fmt.Println("Error while parsing args. ", err)
		return
	}

	app.Start(day, year, workDir)
}

func parse() (int, int, string, error) {
	if len(os.Args) <= 3 {
		return 0, 0, "", errors.New("missing args. 3 minimum")
	}

	flag.Parse()
	year, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		return 0, 0, "", err
	}
	day, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		return 0, 0, "", err
	}
	workDir := flag.Arg(2)

	return day, year, workDir, nil
}
