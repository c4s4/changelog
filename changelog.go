package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
)

type Command func(*Changelog, []string)

var COMMAND_MAPPING = map[string]Command{
	"help":    help,
	"release": release,
	"to":      transform,
}

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
	Rejected   []string
	Notes      []string
}

type Changelog []Release

const (
	ERROR_READING   = 1
	ERROR_PARSING   = 2
	ERROR_RELEASE   = 3
	ERROR_TRANSFORM = 4
	HELP            = `Manage semantic changelog

  changelog                      Print this help screen
  changelog release              Check for release
  changelog release date         Print release date
  changelog release date check   Check that release date wright
  changelog release version      Print release version
  changelog release summary      Print release summary
  changelog to html              Transform changelog to html
  changelog to html stylesheet   Transform to html with a stylesheet
                                 ('style' uses a default stylesheet)
  changelog to markdown          Transform changelog to markdown

The changelog file is searched in current directory. To use a
different changelog, use < character with its path:

  changelog release < path/to/changelog.yml

will check for release a changelog in 'path/to' directory.`
	HELP_COMMAND = "help"
)

var REGEXP_FILENAME = regexp.MustCompile(`^(?i)change(-|_)?log(.yml|.yaml)?$`)

func Error(code int, message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(code)
}

func Errorf(code int, message string, args ...interface{}) {
	Error(code, fmt.Sprintf(message, args...))
}

func readChangelog() []byte {
	stat, err := os.Stdin.Stat()
	if err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
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

func help(changelog *Changelog, args []string) {
	fmt.Println(HELP)
	os.Exit(0)
}

func main() {
	var changelog *Changelog
	var command string
	var args []string
	if len(os.Args) < 2 {
		command = HELP_COMMAND
	} else {
		changelog = parseChangelog(readChangelog())
		command = os.Args[1]
		args = os.Args[2:]
	}
	function := COMMAND_MAPPING[command]
	if function != nil {
		function(changelog, args)
	} else {
		fmt.Printf("Command %s unknown\n", command)
		os.Exit(ERROR_READING)
	}
}
