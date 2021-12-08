//go:build !darwin
// +build !darwin

package input

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type token struct {
	Session string `yaml:"session"`
}

func Token() (session string, err error) {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome[len(configHome)-1] != '/' {
		configHome += "/"
	}
	configDir := fmt.Sprintf("%s%s", configHome, "aocprep")

	tokenConfig := readToken(configDir, "token.yaml")

	return tokenConfig.Session, nil
}

func readToken(dir string, fileName string) *token {
	t := &token{}
	configFile := fmt.Sprintf("%s/%s", dir, fileName)
	fmt.Println(configFile)

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, t)
	if err != nil {
		panic(err)
	}

	return t
}
