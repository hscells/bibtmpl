# bibtmpl

How we generate a publications list for [ielab.io](ielab.io) and [scells.me](scells.me). 

usage:

```
bibtmpl --help                                                               [!?][master]
template BibTex files into HTML
bibtmpl 31.Aug.2018
Usage: bibtmpl --template TEMPLATE --bibtex BIBTEX

Options:
  --template TEMPLATE    path to file to template
  --bibtex BIBTEX        path to BibTex file
  --help, -h             display this help and exit
  --version              display version and exit
```

example:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Example Usage</title>
</head>
<body>
{{ range .Entries }}
<div>
    <span>{{ index .Fields "Author" }}</span>.
    <span>{{index .Fields "Year"}}</span>.
    <span>{{ index .Fields "Title"}}</span>.
    <span>{{ index .Fields "Booktitle"}}</span>
</div>
{{ end }}
</body>
</html>
```
