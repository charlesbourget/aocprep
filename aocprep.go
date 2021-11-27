package main

import (
	"bufio"
	"embed"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

//go:embed resources/template
var f embed.FS

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

	name := fmt.Sprintf("day%d", day)
	path := workDir
	if path[len(path)-1] != '/' {
		path += "/"
	}

	dirPath, err := createDir(path, name)
	if err != nil {
		fmt.Println("Error while creating directory ", err)
		return
	}

	err = createSourceFile(dirPath, name, "resources/template")
	if err != nil {
		fmt.Println("Error while creating source file ", err)
		return
	}
	err = createInputFile(dirPath, "input", day, year)
	if err != nil {
		fmt.Println("Error while creating source file", err)
		return
	}

	fmt.Printf("Structure for day %d created! 🚀\n", day)
}

func createDir(path string, name string) (string, error) {
	if _, err := os.Stat(path + name); err == nil {
		return path + name, nil
	}

	err := os.Mkdir(path+name, os.ModePerm)
	if err != nil {
		return "", err
	}

	return path + name, nil
}

func createSourceFile(dirPath string, name string, src string) error {
	dest := fmt.Sprintf("%s/%s.go", dirPath, name)
	if _, err := os.Stat(dest); err == nil {
		return nil
	}

	source, err := f.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}

func createInputFile(dirPath string, name string, day int, year int) error {
	dest := fmt.Sprintf("%s/%s", dirPath, name)
	if _, err := os.Stat(dest); err == nil {
		return nil
	}

	input, err := FetchInput(day, year)
	if err != nil {
		return err
	}

	file, err := os.Create(dest)
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	_, err = writer.Write(input)
	if err != nil {
		return err
	}

	return nil
}

func parse() (int, int, string, error) {
	if len(os.Args) <= 3 {
		return 0, 0, "", errors.New("missing args. 3 minimum")
	}

	flag.Parse()
	day, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		return 0, 0, "", err
	}
	year, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		return 0, 0, "", err
	}
	workDir := flag.Arg(2)

	return day, year, workDir, nil
}
