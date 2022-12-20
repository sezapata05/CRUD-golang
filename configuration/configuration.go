package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
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
