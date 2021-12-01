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

	err = createSourceFile(dirPath, name, "resources/source.tmpl", day, year)
	if err != nil {
		fmt.Println("Error while creating source file ", err)
		return
	}
	err = createInputFile(dirPath, "input", day, year)
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

func createSourceFile(dirPath string, name string, src string, day int, year int) error {
	dest := fmt.Sprintf("%s/%s.go", dirPath, name)
	if _, err := os.Stat(dest); err == nil {
		return nil
	}

	fileTemplate := FileTemplate()
	t := template.Must(template.New("fileTemplate").Parse(fileTemplate))

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	err = t.Execute(destination, &config{day, year})
	if err != nil {
		panic(err)
	}

	return nil
}

func createInputFile(dirPath string, name string, day int, year int) error {
	dest := fmt.Sprintf("%s/%s", dirPath, name)
	if _, err := os.Stat(dest); err == nil {
		fmt.Println("Input file already exists. Won't download again to prevent overloading the server. 💻")
		return nil
	}

	input, err := input.FetchInput(day, year)
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
