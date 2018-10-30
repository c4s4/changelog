package lib

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
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

// IsPiped tells if content was piped to this process
func IsPiped() bool {
	stat, err := os.Stdin.Stat()
	if err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		return true
	}
	return false
}

// ReadStdin reads standard input and return its content
func ReadStdin() ([]byte, error) {
	source, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("reading changelog from stdin")
	}
	return source, nil
}

// FindChangelog finds source file and return its name
func FindChangelog() (string, error) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return "", fmt.Errorf("could not list current directory")
	}
	for _, file := range files {
		if !file.IsDir() && RegexpFilename.MatchString(file.Name()) {
			return file.Name(), nil
		}
	}
	return "", fmt.Errorf("could not find changelog file")
}

// ReadChangelog reads source file and return contents as array of bytes
func ReadChangelog(file string) ([]byte, error) {
	source, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("reading changelog file '%s'", file)
	}
	return source, nil
}

// ParseChangelog parses source file and return Changelog object
func ParseChangelog(source []byte) (Changelog, error) {
	var changelog Changelog
	err := yaml.Unmarshal(source, &changelog)
	if err != nil {
		return nil, fmt.Errorf("parsing changelog: %s", err.Error())
	}
	return changelog, nil
}

// Help print help and exit
func Help(changelog Changelog, args []string) error {
	fmt.Println(HelpMessage)
	os.Exit(0)
	return nil
}
