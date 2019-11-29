package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"

	"github.com/zbd20/gormin/src/models"
)

var cnf models.Config

func InitConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &cnf)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() models.Config {
	return cnf
}
