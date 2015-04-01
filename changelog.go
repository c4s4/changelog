package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
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
}
var ERROR_READING = 1
var ERROR_PARSING = 2
var ERROR_RELEASE = 3
var REGEXP_DATE = regexp.MustCompile(`^\d\d\d\d-\d\d-\d\d$`)
var REGEXP_VERSION = regexp.MustCompile(`^\d+(\.\d+)?(\.\d+)$`)

func Error(code int, message string) {
	fmt.Println(message)
	os.Exit(code)
}

func Errorf(code int, message string, args ...interface{}) {
	Error(code, fmt.Sprintf(message, args...))
}

func readChangelog() []byte {
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
		}
	}
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
