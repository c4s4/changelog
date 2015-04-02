package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
	"text/template"
)

type Command func(*Changelog, []string)

type Release struct {
	Version    string
	Date       string
	Summary    string
	Added      []string
	Changed    []string
	Deprecated []string
	Removed    []string
	Fixed      []string
	Security   []string
}

type Changelog []Release

var REGEXP_FILENAME = regexp.MustCompile(`^(?i)change(-|_)?log(.yml|.yaml)?$`)
var DEFAULT_COMMAND = "release"
var COMMAND_MAPPING = map[string]Command{
	"release": release,
	"to":      to,
}
var ERROR_READING = 1
var ERROR_PARSING = 2
var ERROR_RELEASE = 3
var ERROR_TO = 4
var REGEXP_DATE = regexp.MustCompile(`^\d\d\d\d-\d\d-\d\d$`)
var REGEXP_VERSION = regexp.MustCompile(`^\d+(\.\d+)?(\.\d+)?$`)
var HTML_TEMPLATE = `<!DOCTYPE html>
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

func Error(code int, message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(code)
}

func Errorf(code int, message string, args ...interface{}) {
	Error(code, fmt.Sprintf(message, args...))
}

func readChangelog() []byte {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// data is being piped to stdin
		source, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			Error(ERROR_READING, "Error reading changelog from stdin")
		}
		return source
	} else {
		// look for changelog in current directory
		files, err := ioutil.ReadDir(".")
		if err != nil {
			Error(ERROR_READING, "Could not list current directory")
		}
		for _, file := range files {
			if !file.IsDir() && REGEXP_FILENAME.MatchString(file.Name()) {
				source, err := ioutil.ReadFile(file.Name())
				if err != nil {
					Errorf(ERROR_READING, "Error reading changelog file '%s'\n", file.Name())
				}
				return source
			}
		}
		Error(ERROR_READING, "No changelog file found")
		return []byte{}
	}
}

func parseChangelog(source []byte) *Changelog {
	var changelog Changelog
	err := yaml.Unmarshal(source, &changelog)
	if err != nil {
		Errorf(ERROR_PARSING, "Error parsing changelog: %s\n", err.Error())
	}
	return &changelog
}

////////////////////////////////////////////////////////////////////////////////
//                                  COMMANDS                                  //
////////////////////////////////////////////////////////////////////////////////

func checkRelease(release Release) {
	if release.Version == "" {
		Error(ERROR_RELEASE, "Release version is empty")
	}
	if !REGEXP_VERSION.MatchString(release.Version) {
		Errorf(ERROR_RELEASE, "Release version '%s' is not a valid semantic version number", release.Version)
	}
	if release.Date == "" {
		Error(ERROR_RELEASE, "Release date is empty")
	}
	if !REGEXP_DATE.MatchString(release.Date) {
		Errorf(ERROR_RELEASE, "Release date '%s' is not valid ISO format", release.Date)
	}
	if release.Summary == "" {
		Error(ERROR_RELEASE, "Release summary is empty")
	}
}

func checkChangelog(changelog *Changelog) {
	if len(*changelog) == 0 {
		Error(ERROR_RELEASE, "Release is empy")
	}
	for _, release := range *changelog {
		checkRelease(release)
	}
}

func release(changelog *Changelog, args []string) {
	checkChangelog(changelog)
	if len(args) > 0 {
		if args[0] == "summary" {
			fmt.Println((*changelog)[0].Summary)
		} else if args[0] == "date" {
			fmt.Println((*changelog)[0].Date)
		} else if args[0] == "version" {
			fmt.Println((*changelog)[0].Version)
		} else {
			Errorf(ERROR_RELEASE, "Unknown release argument %s", args[0])
		}
	}
}

type HtmlTemplateData struct {
	Changelog   *Changelog
	Stylesheets []string
}

func toHtml(changelog *Changelog, args []string) {
	stylesheets := make([]string, 0)
	for _, file := range args {
		stylesheet, err := ioutil.ReadFile(file)
		if err != nil {
			Errorf(ERROR_TO, "Error loading stylesheet %s: %s", file, err.Error())
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
		Errorf(ERROR_TO, "Error processing template: %s", err)
	}
}

func to(changelog *Changelog, args []string) {
	checkChangelog(changelog)
	if len(args) < 1 {
		Error(ERROR_TO, "You must pass format to convert to")
	}
	format := args[0]
	if format != "html" {
		Errorf(ERROR_TO, "Unknown format %s", args[0])
	}
	toHtml(changelog, args[1:])
}

func main() {
	changelog := parseChangelog(readChangelog())
	var command string
	var args []string
	if len(os.Args) < 2 {
		command = DEFAULT_COMMAND
		args = []string(nil)
	} else {
		command = os.Args[1]
		args = os.Args[2:]
	}
	function := COMMAND_MAPPING[command]
	if function != nil {
		function(changelog, args)
	} else {
		fmt.Printf("Command %s unknown\n", command)
		os.Exit(3)
	}
}
