package main

import (
	"github.com/alexflint/go-arg"
	"os"
	"github.com/hscells/bibtmpl"
)

type args struct {
	Template string `arg:"required,help:path to file to template"`
	BibTex   string `arg:"required,help:path to BibTex file"`
}

func (args) Version() string {
	return "bibtmpl 31.Aug.2018"
}

func (args) Description() string {
	return `template BibTex files into HTML`
}

func main() {
	// Parse the args into the struct.
	var args args
	arg.MustParse(&args)

	output := os.Stdout

	// Template the file.
	err := bibtmpl.Template(args.Template, args.BibTex, output)

	if err != nil {
		panic(err)
	}
}
