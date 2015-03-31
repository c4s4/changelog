package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
)

type Command func(*ChangeLog, []string)

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

type ChangeLog []Release

var REGEXP_FILENAME = regexp.MustCompile("^(?i)change(-|_)?log(.yml|.yaml)?$")
var DEFAULT_COMMAND = "release"
var COMMAND_MAPPING = map[string]Command{
	"release": release,
}
var ERROR_READING = 1
var ERROR_PARSING = 2
var ERROR_RELEASE = 3

func readChangeLog() []byte {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Could not list current directory")
		os.Exit(ERROR_READING)
	}
	for _, file := range files {
		if !file.IsDir() && REGEXP_FILENAME.MatchString(file.Name()) {
			source, err := ioutil.ReadFile(file.Name())
			if err != nil {
				fmt.Printf("Error reading changelog file '%s'\n", file.Name())
				os.Exit(ERROR_READING)
			}
			return source
		}
	}
	fmt.Println("No changelog file found")
	os.Exit(ERROR_READING)
	return []byte{}
}

func parseChangeLog(source []byte) *ChangeLog {
	var changelog ChangeLog
	err := yaml.Unmarshal(source, &changelog)
	if err != nil {
		fmt.Printf("Error parsing changelog: %s\n", err)
		os.Exit(ERROR_PARSING)
	}
	return &changelog
}

//////////////////////////////////////////////
//                 COMMANDS                 //
//////////////////////////////////////////////

func checkRelease(changelog *ChangeLog) {
	if len(*changelog) == 0 {
		fmt.Println("Release is empy")
		os.Exit(ERROR_RELEASE)
	}
	release := (*changelog)[0]
	if release.Version == "" {
		fmt.Println("Release version is empty")
		os.Exit(ERROR_RELEASE)
	}
	if release.Date == "" {
		fmt.Println("Release date is empty")
		os.Exit(ERROR_RELEASE)
	}
	if release.Summary == "" {
		fmt.Println("Release summary is empty")
		os.Exit(ERROR_RELEASE)
	}
}

func release(changelog *ChangeLog, args []string) {
	checkRelease(changelog)
	if len(args) > 0 {
		if args[0] == "summary" {
			fmt.Println((*changelog)[0].Summary)
		} else if args[0] == "date" {
			fmt.Println((*changelog)[0].Date)
		}
	}
}

func main() {
	changelog := parseChangeLog(readChangeLog())
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
