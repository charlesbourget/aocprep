package main

import (
	"flag"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/charlesbourget/aocprep/app"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: prepare <day> <work dir>")
		flag.PrintDefaults()
	}
	parser := argparse.NewParser("aocprep", "Prepare for an Advent of Code day")
	year := parser.Int("y", "year", &argparse.Options{Required: true, Help: "Year"})
	day := parser.Int("d", "day", &argparse.Options{Required: true, Help: "Day"})
	workDir := parser.String("w", "workDir", &argparse.Options{Required: false, Help: "Workdir"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	if *workDir == "" {
		*workDir = "."
	}

	app.Start(*day, *year, *workDir)
}
