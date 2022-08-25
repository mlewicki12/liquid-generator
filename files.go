package liquidgenerator

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type YamlData = map[string]interface{}

func read(path string) YamlData {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	data := make(map[string]interface{})
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Fatalln(err)
	}

	return data
}
