package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	mlk := sage{
		Name:  "MLK",
		Motto: "I am a communist",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	sages := []sage{buddha, gandhi, mlk, jesus}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
