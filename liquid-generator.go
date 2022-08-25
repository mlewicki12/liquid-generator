package liquidgenerator

import (
	"fmt"
	"log"

	liquid "github.com/osteele/liquid"
)

func Print() {
	fmt.Printf("hello world\n")
}

func GenerateTemplate() {
	engine := liquid.NewEngine()
	template := `<h1>{{ page.title }}</h1>`
	bindings := map[string]interface{}{
		"page": map[string]interface{}{
			"title": "Introduction",
		},
	}

	out, err := engine.ParseAndRenderString(template, bindings)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(out)
}
