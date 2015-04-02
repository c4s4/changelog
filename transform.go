package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

const (
	HTML_TEMPLATE = `<!DOCTYPE html>
<html>
<head>
<title>Change Log</title>
<meta charset="utf-8">
{{ range $stylesheet := .Stylesheets }}
<style type="text/css">
{{ $stylesheet }}
</style>
{{ end }}
</head>
<body>
<h1>Change Log</h1>
{{ range $release := .Changelog }}
<h2>Release {{ .Version }} ({{ .Date }})</h2>
<p>{{ .Summary }}</p>
{{ if .Added }}
<h3>Added</h3>
<ul>
{{ range $entry := .Added }}
<li>{{ . }}</li>
{{ end }}
</ul>
{{ end }}
{{ if .Changed }}
<h3>Changed</h3>
<ul>
{{ range $entry := .Changed }}
<li>{{ . }}</li>
{{ end }}
</ul>
{{ end }}
{{ if .Deprecated }}
<h3>Deprecated</h3>
<ul>
{{ range $entry := .Deprecated }}
<li>{{ . }}</li>
{{ end }}
</ul>
{{ end }}
{{ if .Removed }}
<h3>Removed</h3>
<ul>
{{ range $entry := .Removed }}
<li>{{ . }}</li>
{{ end }}
</ul>
{{ end }}
{{ if .Fixed }}
<h3>Fixed</h3>
<ul>
{{ range $entry := .Fixed }}
<li>{{ . }}</li>
{{ end }}
</ul>
{{ end }}
{{ if .Security }}
<h3>Security</h3>
<ul>
{{ range $entry := .Security }}
<li>{{ . }}</li>
{{ end }}
</ul>
{{ end }}
{{ end }}
</body>
</html>`
)

type HtmlTemplateData struct {
	Changelog   *Changelog
	Stylesheets []string
}

func toHtml(changelog *Changelog, args []string) {
	stylesheets := make([]string, 0)
	for _, file := range args {
		stylesheet, err := ioutil.ReadFile(file)
		if err != nil {
			Errorf(ERROR_TRANSFORM, "Error loading stylesheet %s: %s", file, err.Error())
		}
		stylesheets = append(stylesheets, string(stylesheet))
	}
	data := HtmlTemplateData{
		Stylesheets: stylesheets,
		Changelog:   changelog,
	}
	t := template.Must(template.New("changelog").Parse(HTML_TEMPLATE))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		Errorf(ERROR_TRANSFORM, "Error processing template: %s", err)
	}
}

func transform(changelog *Changelog, args []string) {
	checkChangelog(changelog)
	if len(args) < 1 {
		Error(ERROR_TRANSFORM, "You must pass format to transform to")
	}
	format := args[0]
	if format != "html" {
		Errorf(ERROR_TRANSFORM, "Unknown format %s", args[0])
	}
	toHtml(changelog, args[1:])
}
