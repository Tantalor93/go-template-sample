package main

import (
	"fmt"
	"os"
	"text/template"
)

type Holder struct {
	Name string
	Value int
	Names []string
}

func main()  {
	holder := Holder{"Petr", 1, []string{"alfa", "beta", "charlie"}}

	templateWithCondition(holder)
	fmt.Println()
	templateWithRange(holder)
}


func templateWithRange(holder Holder) {
	tmpl := template.New("benky template")
	tmpl, error := tmpl.Parse("[{{ range .Names }} {{ . }} {{end}}]")
	if error != nil {
		panic("Unable to parse template: " + error.Error())
	}
	tmpl.Execute(os.Stdout, holder)
}

func templateWithCondition(holder Holder) {
	tmpl := template.New("benky template")
	tmpl, error := tmpl.Parse("Hello '{{.Name}}' and {{if eq .Value 0 }} value is zero! {{else}} value is non zero! {{end}}")
	if error != nil {
		panic("Unable to parse template: " + error.Error())
	}
	tmpl.Execute(os.Stdout, holder)
}