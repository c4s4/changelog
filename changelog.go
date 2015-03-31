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
var DEFAULT_COMMAND = "NO_COMMAND"
var COMMAND_MAPPING = map[string]Command {
  "summary": summary,
}

func readChangeLog() []byte {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Could not list current directory")
		os.Exit(1)
	}
	for _, file := range files {
		if !file.IsDir() && REGEXP_FILENAME.MatchString(file.Name()) {
			source, err := ioutil.ReadFile(file.Name())
			if err != nil {
				fmt.Printf("Error reading changelog file '%s'\n", file.Name())
				os.Exit(1)
			}
			return source
		}
	}
	fmt.Println("No changelog file found")
	os.Exit(2)
	return []byte{}
}

func parseChangeLog(source []byte) *ChangeLog {
	var changelog ChangeLog
	err := yaml.Unmarshal(source, &changelog)
	if err != nil {
		fmt.Printf("Error parsing changelog: %s\n", err)
		os.Exit(3)
	}
	return &changelog
}

func summary(changelog *ChangeLog, args []string) {
  fmt.Println(*changelog[0].Summary)
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
	  fmt.Printf("Command %s unknown", command)
	  os.Exit(3)
	}
}
