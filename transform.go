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
<title>Changelog</title>
<meta charset="utf-8">
{{ range $stylesheet := .Stylesheets }}
<style type="text/css">
{{ $stylesheet }}
</style>
{{ end }}
</head>
<body>
<h1>Changelog</h1>
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
{{ if .Rejected }}
<h3>Rejected</h3>
<ul>
{{ range $entry := .Rejected }}
<li>{{ . }}</li>
{{ end }}
</ul>
{{ end }}
{{ if .Notes }}
<h3>Notes</h3>
<ul>
{{ range $entry := .Notes }}
<li>{{ . }}</li>
{{ end }}
</ul>
{{ end }}
{{ end }}
</body>
</html>`
	STYLESHEET = `
body {
  font-family: Helvetica, arial, sans-serif;
  font-size: 16px;
  line-height: 1.4;
  padding-top: 10px;
  padding-bottom: 10px;
  background-color: white;
  padding: 30px;
  color: #333;
}

body > *:first-child {
  margin-top: 0 !important;
}

body > *:last-child {
  margin-bottom: 0 !important;
}

a {
  color: #4183C4;
  text-decoration: none;
}

a.absent {
  color: #cc0000;
}

a.anchor {
  display: block;
  padding-left: 30px;
  margin-left: -30px;
  cursor: pointer;
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
}

h1, h2, h3, h4, h5, h6 {
  margin: 20px 0 10px;
  padding: 0;
  font-weight: bold;
  -webkit-font-smoothing: antialiased;
  cursor: text;
  position: relative;
}

h2:first-child, h1:first-child, h1:first-child + h2, h3:first-child, h4:first-child, h5:first-child, h6:first-child {
  margin-top: 0;
  padding-top: 0;
}

h1:hover a.anchor, h2:hover a.anchor, h3:hover a.anchor, h4:hover a.anchor, h5:hover a.anchor, h6:hover a.anchor {
  text-decoration: none;
}

h1 tt, h1 code {
  font-size: inherit;
}

h2 tt, h2 code {
  font-size: inherit;
}

h3 tt, h3 code {
  font-size: inherit;
}

h4 tt, h4 code {
  font-size: inherit;
}

h5 tt, h5 code {
  font-size: inherit;
}

h6 tt, h6 code {
  font-size: inherit;
}

h1 {
  font-size: 32px;
  color: black;
}

h2 {
  font-size: 28px;
  border-bottom: 1px solid #cccccc;
  color: black;
}

h3 {
  font-size: 20px;
}

h4 {
  font-size: 16px;
}

h5 {
  font-size: 14px;
}

h6 {
  color: #777777;
  font-size: 14px;
}

p, blockquote, ul, ol, dl, table, pre {
  margin: 15px 0;
}

hr {
  background: transparent url("http://tinyurl.com/bq5kskr") repeat-x 0 0;
  border: 0 none;
  color: #cccccc;
  height: 4px;
  padding: 0;
}

body > h2:first-child {
  margin-top: 0;
  padding-top: 0;
}

body > h1:first-child {
  margin-top: 0;
  padding-top: 0;
}

body > h1:first-child + h2 {
  margin-top: 0;
  padding-top: 0;
}

body > h3:first-child, body > h4:first-child, body > h5:first-child, body > h6:first-child {
  margin-top: 0;
  padding-top: 0;
}

a:first-child h1, a:first-child h2, a:first-child h3, a:first-child h4, a:first-child h5, a:first-child h6 {
  margin-top: 0;
  padding-top: 0;
}

h1 p, h2 p, h3 p, h4 p, h5 p, h6 p {
  margin-top: 0;
}

li p.first {
  display: inline-block;
}

ul, ol {
  padding-left: 30px;
}

ul :first-child, ol :first-child {
  margin-top: 0;
}

ul :last-child, ol :last-child {
  margin-bottom: 0;
}

dl {
  padding: 0;
}

dl dt {
  font-size: 14px;
  font-weight: bold;
  font-style: italic;
  padding: 0;
  margin: 15px 0 5px;
}

dl dt:first-child {
  padding: 0;
}

dl dt > :first-child {
  margin-top: 0;
}

dl dt > :last-child {
  margin-bottom: 0;
}

dl dd {
  margin: 0 0 15px;
  padding: 0 15px;
}

dl dd > :first-child {
  margin-top: 0;
}

dl dd > :last-child {
  margin-bottom: 0;
}

blockquote {
  border-left: 4px solid #dddddd;
  padding: 0 15px;
  color: #777777;
}

blockquote > :first-child {
  margin-top: 0;
}

blockquote > :last-child {
  margin-bottom: 0;
}

table {
  padding: 0;
  border-spacing: 0;
  border-collapse: collapse;
}

table tr {
  border-top: 1px solid #cccccc;
  background-color: white;
  margin: 0;
  padding: 0;
}

table tr:nth-child(2n) {
  background-color: #f7f7f7;
}

table th {
  border: 1px solid #cccccc;
  font-weight: bold;
  text-align: left;
  margin: 0;
  padding: 5px 10px;
  background-color: #f0f0f0;
}

table td {
  border: 1px solid #cccccc;
  text-align: left;
  margin: 0;
  padding: 5px 10px;
}

table tr th :first-child, table tr td :first-child {
  margin-top: 0;
}

table tr th :last-child, table tr td :last-child {
  margin-bottom: 0;
}

img {
  max-width: 100%;
}

span.frame {
  display: block;
  overflow: hidden;
}

span.frame > span {
  border: 1px solid #dddddd;
  display: block;
  float: left;
  overflow: hidden;
  margin: 13px 0 0;
  padding: 7px;
  width: auto;
}

span.frame span img {
  display: block;
  float: left;
}

span.frame span span {
  clear: both;
  color: #333333;
  display: block;
  padding: 5px 0 0;
}

span.align-center {
  display: block;
  overflow: hidden;
  clear: both;
}

span.align-center > span {
  display: block;
  overflow: hidden;
  margin: 13px auto 0;
  text-align: center;
}

span.align-center span img {
  margin: 0 auto;
  text-align: center;
}

span.align-right {
  display: block;
  overflow: hidden;
  clear: both;
}

span.align-right > span {
  display: block;
  overflow: hidden;
  margin: 13px 0 0;
  text-align: right;
}

span.align-right span img {
  margin: 0;
  text-align: right;
}

span.float-left {
  display: block;
  margin-right: 13px;
  overflow: hidden;
  float: left;
}

span.float-left span {
  margin: 13px 0 0;
}

span.float-right {
  display: block;
  margin-left: 13px;
  overflow: hidden;
  float: right;
}

span.float-right > span {
  display: block;
  overflow: hidden;
  margin: 13px auto 0;
  text-align: right;
}

code, tt {
  margin: 0 2px;
  padding: 0 5px;
  white-space: nowrap;
  border: 1px solid #eaeaea;
  background-color: #f8f8f8;
  border-radius: 3px;
}

pre code {
  margin: 0;
  padding: 0;
  white-space: pre;
  border: none;
  background: transparent;
}

.highlight pre {
  background-color: #f8f8f8;
  border: 1px solid #cccccc;
  font-size: 13px;
  line-height: 19px;
  overflow: auto;
  padding: 6px 10px;
  border-radius: 3px;
}

pre {
  background-color: #f8f8f8;
  border: 1px solid #cccccc;
  font-size: 13px;
  line-height: 19px;
  overflow: auto;
  padding: 6px 10px;
  border-radius: 3px;
}

pre code, pre tt {
  background-color: transparent;
  border: none;
}`

	MD_TEMPLATE = `# Changelog

{{ range $release := .Changelog }}## Release {{ .Version }} ({{ .Date }})

{{ if .Summary }}{{ .Summary }}{{ end }}

{{ if .Added }}### Added

{{ range $entry := .Added }}- {{ . }}
{{ end }}{{ end }}{{ if .Changed }}
### Changed

{{ range $entry := .Changed }}- {{ . }}
{{ end }}{{ end }}{{ if .Deprecated }}
### Deprecated

{{ range $entry := .Deprecated }}- {{ . }}
{{ end }}{{ end }}{{ if .Removed }}
### Removed

{{ range $entry := .Removed }}- {{ . }}
{{ end }}{{ end }}{{ if .Fixed }}
### Fixed

{{ range $entry := .Fixed }}- {{ . }}
{{ end }}{{ end }}{{ if .Security }}
### Security

{{ range $entry := .Security }}- {{ . }}
{{ end }}{{ end }}{{ if .Rejected }}
### Rejected

{{ range $entry := .Rejected }}- {{ . }}
{{ end }}{{ end }}{{ if .Notes }}
### Notes

{{ range $entry := .Notes }}- {{ . }}
{{ end }}{{ end }}
{{ end }}`

	MD_TEMPLATE_RELEASE = `# Release {{ .Version }} ({{ .Date }})

{{ if .Summary }}{{ .Summary }}{{ end }}

{{ if .Added }}## Added

{{ range $entry := .Added }}- {{ . }}
{{ end }}{{ end }}{{ if .Changed }}
## Changed

{{ range $entry := .Changed }}- {{ . }}
{{ end }}{{ end }}{{ if .Deprecated }}
## Deprecated

{{ range $entry := .Deprecated }}- {{ . }}
{{ end }}{{ end }}{{ if .Removed }}
## Removed

{{ range $entry := .Removed }}- {{ . }}
{{ end }}{{ end }}{{ if .Fixed }}
## Fixed

{{ range $entry := .Fixed }}- {{ . }}
{{ end }}{{ end }}{{ if .Security }}
## Security

{{ range $entry := .Security }}- {{ . }}
{{ end }}{{ end }}{{ if .Rejected }}
## Rejected

{{ range $entry := .Rejected }}- {{ . }}
{{ end }}{{ end }}{{ if .Notes }}
## Notes

{{ range $entry := .Notes }}- {{ . }}
{{ end }}{{ end }}`
)

type TemplateDataChangelog struct {
	Changelog   Changelog
	Stylesheets []string
}

func toHtml(changelog Changelog, args []string) {
	stylesheets := make([]string, 0)
	for _, file := range args {
		var stylesheet []byte
		var err error
		if file == "style" {
			stylesheet = []byte(STYLESHEET)
		} else {
			stylesheet, err = ioutil.ReadFile(file)
			if err != nil {
				Errorf(ERROR_TRANSFORM, "Error loading stylesheet %s: %s", file, err.Error())
			}
		}
		stylesheets = append(stylesheets, string(stylesheet))
	}
	data := TemplateDataChangelog{
		Stylesheets: stylesheets,
		Changelog:   changelog,
	}
	t := template.Must(template.New("changelog").Parse(HTML_TEMPLATE))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		Errorf(ERROR_TRANSFORM, "Error processing template: %s", err)
	}
}

func toMarkdown(changelog Changelog) {
	data := TemplateDataChangelog{
		Stylesheets: nil,
		Changelog:   changelog,
	}
	t := template.Must(template.New("changelog").Parse(MD_TEMPLATE))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		Errorf(ERROR_TRANSFORM, "Error processing template: %s", err)
	}
}

func releaseToMarkdown(release Release) {
	t := template.Must(template.New("release").Parse(MD_TEMPLATE_RELEASE))
	err := t.Execute(os.Stdout, release)
	if err != nil {
		Errorf(ERROR_TRANSFORM, "Error processing template: %s", err)
	}
}

func transform(changelog Changelog, args []string) {
	checkChangelog(changelog)
	if len(args) < 1 {
		Error(ERROR_TRANSFORM, "You must pass format to transform to")
	}
	format := args[0]
	if format == "html" {
		toHtml(changelog, args[1:])
	} else if format == "markdown" {
		toMarkdown(changelog)
	} else {
		Errorf(ERROR_TRANSFORM, "Unknown format %s", args[0])
	}
}
