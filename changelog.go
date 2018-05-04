package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Command is a changelog command implemented with a function
type Command func(Changelog, []string)

// CommandMapping maps command names with command functions
var CommandMapping = map[string]Command{
	"Help":    Help,
	"release": release,
	"to":      transform,
}

// Release contains information about a release
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

// Changelog is a list of releases
type Changelog []Release

const (
	// ErrorReading denotes an error reading changelog file
	ErrorReading = 1
	// ErrorParsing denotes an error parsing changelog file
	ErrorParsing = 2
	// ErrorRelease denotes an error in a release
	ErrorRelease = 3
	// ErrorTransform denotes an error transforming changelog
	ErrorTransform = 4
	// HelpMessage is the help text
	HelpMessage = `Manage semantic changelog

  changelog                      Print this Help screen
  changelog release              Check for release
  changelog release date         Print release date
  changelog release date check   Check that release date wright
  changelog release version      Print release version
  changelog release summary      Print release summary
  changelog release to markdown  Print release changelog in markdown
  changelog to html              Transform changelog to html
  changelog to html stylesheet   Transform to html with a stylesheet
                                 ('style' uses a default stylesheet)
  changelog to markdown          Transform changelog to markdown

You can add 'next' after changelog command to consider next to last release
instead of the last, or '-N' go go back in past Nth release.

The changelog file is searched in current directory. To use a different
changelog, use < character with its path:

  changelog release < path/to/changelog.yml

will check for release a changelog in 'path/to' directory.`
	// HelpCommand is the command for help
	HelpCommand = "Help"
)

// RegexpFilename is the regular expression for changelog filename
var RegexpFilename = regexp.MustCompile(`^(?i)change(-|_)?log(.yml|.yaml)?$`)

// Error print an error message and exit with an error code
func Error(code int, message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(code)
}

// Errorf print an error message with arguments and exit with an error code
func Errorf(code int, message string, args ...interface{}) {
	Error(code, fmt.Sprintf(message, args...))
}

func readChangelog() []byte {
	stat, err := os.Stdin.Stat()
	if err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		// data is being piped to stdin
		source, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			Error(ErrorReading, "Error reading changelog from stdin")
		}
		return source
	}
	// look for changelog in current directory
	files, err := ioutil.ReadDir(".")
	if err != nil {
		Error(ErrorReading, "Could not list current directory")
	}
	for _, file := range files {
		if !file.IsDir() && RegexpFilename.MatchString(file.Name()) {
			source, err := ioutil.ReadFile(file.Name())
			if err != nil {
				Errorf(ErrorReading, "Error reading changelog file '%s'\n", file.Name())
			}
			return source
		}
	}
	Error(ErrorReading, "No changelog file found")
	return []byte{}
}

func parseChangelog(source []byte) Changelog {
	var changelog Changelog
	err := yaml.Unmarshal(source, &changelog)
	if err != nil {
		Errorf(ErrorParsing, "Error parsing changelog: %s\n", err.Error())
	}
	return changelog
}

// Help print help and exit
func Help(changelog Changelog, args []string) {
	fmt.Println(HelpMessage)
	os.Exit(0)
}

func main() {
	var changelog Changelog
	var command string
	var args []string
	if len(os.Args) < 2 {
		command = HelpCommand
	} else {
		changelog = parseChangelog(readChangelog())
		command = os.Args[1]
		if command == "next" {
			command = "-1"
		}
		if strings.HasPrefix(command, "-") {
			delta, err := strconv.Atoi(command[1:])
			if err != nil || delta >= len(changelog) {
				fmt.Printf("Bad shift '%s'\n", command)
				os.Exit(1)
			}
			changelog = changelog[delta:]
			command = os.Args[2]
			args = os.Args[3:]
		} else {
			args = os.Args[2:]
		}
	}
	function := CommandMapping[command]
	if function != nil {
		function(changelog, args)
	} else {
		fmt.Printf("Command '%s' unknown\n", command)
		os.Exit(ErrorReading)
	}
}
