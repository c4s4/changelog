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
type Command func(Changelog, []string) error

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

func isPiped() bool {
	stat, err := os.Stdin.Stat()
	if err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		return true
	}
	return false
}

func findChangelog() (string, error) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return "", fmt.Errorf("Could not list current directory")
	}
	for _, file := range files {
		if !file.IsDir() && RegexpFilename.MatchString(file.Name()) {
			return file.Name(), nil
		}
	}
	return "", nil
}

func readChangelog(file string) ([]byte, error) {
	// file was passed
	if file != "" {
		source, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("Error reading changelog file '%s'", file)
		}
		return source, nil
	}
	// look for changelog piped in stdin
	if isPiped() {
		source, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("Error reading changelog from stdin")
		}
		return source, nil
	}
	// look for changelog in current directory
	file, err := findChangelog()
	if err != nil {
		return nil, err
	}
	if file != "" {
		source, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("Error reading changelog file '%s'", file)
		}
		return source, nil
	}
	return nil, fmt.Errorf("No changelog file found")
}

func parseChangelog(source []byte) (Changelog, error) {
	var changelog Changelog
	err := yaml.Unmarshal(source, &changelog)
	if err != nil {
		return nil, fmt.Errorf("Error parsing changelog: %s", err.Error())
	}
	return changelog, nil
}

// Help print help and exit
func Help(changelog Changelog, args []string) error {
	fmt.Println(HelpMessage)
	os.Exit(0)
	return nil
}

func main() {
	var changelog Changelog
	var command string
	var args []string
	if len(os.Args) < 2 {
		command = HelpCommand
	} else {
		source, err := readChangelog("")
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s", err.Error())
		}
		changelog, err = parseChangelog(source)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s", err.Error())
		}
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
