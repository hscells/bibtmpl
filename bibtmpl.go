package bibtmpl

import (
	"bytes"
	"fmt"
	"github.com/nickng/bibtex"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

const fieldYear = "Year"
const fieldUrl = "Url"

// Parse reads a file and parses it into a BibTex record.
func Parse(filename string) (*bibtex.BibTex, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// The library that parses the files can't handle comments.
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// So here we remove the comments.
	buff := bytes.NewBuffer(b)
	p := bytes.NewBufferString("")
	for _, line := range strings.Split(buff.String(), "\n") {
		if len(line) > 0 {
			if line[0] == '%' {
				continue
			}
		}
		p.WriteString(line)
	}

	parsed, err := bibtex.Parse(p)
	if err != nil {
		return nil, err
	}

	// Ensure each entry has a `Year`, and warn if if does not.
	for _, entry := range parsed.Entries {
		if _, ok := entry.Fields[fieldYear]; !ok {
			fmt.Printf("[WARN] entry %s does not have a `Year` field (will render incorrectly)\n", entry.CiteName)
			entry.Fields[fieldYear] = bibtex.NewBibConst("0000")
		}
	}

	sort.Slice(parsed.Entries, func(i, j int) bool {
		x, _ := strconv.Atoi(parsed.Entries[i].Fields[fieldYear].String())
		y, _ := strconv.Atoi(parsed.Entries[j].Fields[fieldYear].String())
		return x > y
	})

	return parsed, nil
}

// Template a file with a BibTex file and writes the result to an output file.
func Template(t, b string, o io.Writer) error {
	bf, err := Parse(b)
	if err != nil {
		return err
	}

	tf, err := os.OpenFile(t, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer tf.Close()

	text, err := ioutil.ReadAll(tf)
	if err != nil {
		return err
	}

	tmpl, err := template.New(b).Parse(string(text))
	if err != nil {
		return err
	}

	return tmpl.Execute(o, bf)
}
