package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/joho/godotenv"
)

type DataBase struct {
	Typedb string `yaml:"type_db"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
}

type Configuration struct {
	DB DataBase `yaml:"Database"`
}

func LoadConfig(filename string) (*Configuration, error) {
	configdb := &Configuration{}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, configdb)
	if err != nil {
		return nil, err
	}

	return configdb, nil
}

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
