package main

import (
	"html/template"
	"os"
)

func learnCompositeTemplate() {

	data := map[string]string{
		"Test":         "Whatever",
		"Another test": "Haha",
		"Final test":   "whatever",
	}

	t, _ := template.ParseFiles("templates/compositeTemplate.html")

	t.Execute(os.Stdout, data)
}
