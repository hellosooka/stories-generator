package main

import (
	"github.com/hellosooka/vue-stories-generator/src/files"
	"log"
	"os"
	"text/template"
)

type Test struct {
	bebra string
}

func main() {
	bebra := Test{bebra: "test"}
	temps, err := fileparser.ParseTemplatePaths("./templates/")
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("vue-story").ParseFiles(temps...)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, bebra)
	if err != nil {
		panic(err)
	}

}

// Vue/utils
