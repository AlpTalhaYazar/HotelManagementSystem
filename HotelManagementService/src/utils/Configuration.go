package utils

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Databases []Database `json:"databases"`
}

type Database struct {
	Name             string `json:"name"`
	ConnectionString string `json:"connectionString"`
}

func ReadConfiguration() Configuration {
	projectRootPath, err := os.Getwd()

	file, err := os.Open(projectRootPath + "/conf.json")

	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)

	if err != nil {
		panic(err)
	}

	return configuration
}
