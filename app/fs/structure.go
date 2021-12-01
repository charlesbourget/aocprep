package fs

import (
	"bufio"
	"fmt"
	"github.com/charlesbourget/aocprep/app/input"
	"os"
	"text/template"
)

type config struct {
	Day  int
	Year int
}

func CreateStructure(path string, name string, year int, day int) {
	dirPath, err := createDir(path, name)
	if err != nil {
		fmt.Println("Error while creating directory ", err)
		return
	}

	// Source file
	err = createFileTemplate(dirPath, name, day, year, SourceFileTemplate())
	if err != nil {
		fmt.Println("Error while creating source file ", err)
		return
	}
	// Test source file
	err = createFileTemplate(dirPath, name+"_test", day, year, TestFileTemplate())
	if err != nil {
		fmt.Println("Error while creating source file ", err)
		return
	}
	// Input file
	err = createInputFile(dirPath, "input", day, year)
	if err != nil {
		fmt.Println("Error while creating source file", err)
		return
	}
	// Test input file
	_, err = createFile(dirPath, "input", "test")
	if err != nil {
		fmt.Println("Error while creating source file", err)
		return
	}
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

func createFileTemplate(dirPath string, name string, day int, year int, fileTemplate string) error {
	file, err := createFile(dirPath, name, "go")

	t := template.Must(template.New("fileTemplate").Parse(fileTemplate))

	defer file.Close()

	err = t.Execute(file, &config{day, year})
	if err != nil {
		return err
	}

	return nil
}

func createInputFile(dirPath string, name string, day int, year int) error {
	file, err := createFile(dirPath, name, "")
	if err != nil {
		return err
	}
	if file == nil {
		fmt.Println("Input already exists. Skipping")
		return nil
	}

	input, err := input.FetchInput(day, year)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	_, err = writer.Write(input)
	if err != nil {
		return err
	}

	return nil
}

func createFile(dirPath string, name string, ext string) (*os.File, error) {
	dest := fmt.Sprintf("%s/%s", dirPath, name)
	if ext != "" {
		dest = fmt.Sprintf("%s.%s", dest, ext)
	}
	if _, err := os.Stat(dest); err == nil {
		// File already exists
		return &os.File{}, nil
	}

	file, err := os.Create(dest)
	if err != nil {
		return &os.File{}, err
	}

	return file, nil
}
