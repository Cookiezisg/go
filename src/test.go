package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!")

	println(t.Tree.Root.String())
	p := Person{UserName: "Weilin Sun"}
	t.Execute(os.Stdout, p)
}
