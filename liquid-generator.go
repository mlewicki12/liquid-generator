package liquidgenerator

import (
	"fmt"
	"log"
	"os"

	liquid "github.com/osteele/liquid"
)

var engine *liquid.Engine = liquid.NewEngine()

type Variables = liquid.Bindings

// GenerateTemplate takes a path to a Liquid template and variables and generates a Component out of it
func GenerateTemplate(path string, bindings Variables) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error generating a template\n%v", err)
	}

	out, err := engine.ParseAndRender(data, bindings)
	if err != nil {
		log.Fatalln(err)
	}

	return out, nil
}
