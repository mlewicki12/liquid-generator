package liquidgenerator

import (
	"fmt"
	"log"

	liquid "github.com/osteele/liquid"
)

var engine *liquid.Engine = liquid.NewEngine()

func Print() {
	fmt.Printf("hello world\n")
}

func GenerateTemplate() {
	data := read("resources/themes/pink/generate.yaml")
	for k, v := range data {
		fmt.Printf("[%v]: %v\n", k, v)
	}

	template := `<h1>{{ page.title }}</h1>`
	bindings := liquid.Bindings{
		"page": liquid.Bindings{
			"title": "Introduction",
		},
	}

	out, err := engine.ParseAndRenderString(template, bindings)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(out)
}
