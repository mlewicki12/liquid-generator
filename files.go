package liquidgenerator

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type YamlData = map[string]interface{}

func readMap(path string) YamlData {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	data := make(YamlData)
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Fatalln(err)
	}

	return data
}

func readSlice(path string) []interface{} {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	data := make([]interface{}, 2, 24)
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Fatalln(err)
	}

	return data
}
