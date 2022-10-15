package main

import (
	"os"
	"text/template"
)

func main() {
	template.Must(template.New("test").Parse("{{.}}!")).Execute(os.Stdout, "Hello, playground")
}
