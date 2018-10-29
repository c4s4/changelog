package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"github.com/c4s4/changelog"
)

func printError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	var changelog Changelog
	var command string
	var args []string
	if len(os.Args) < 2 {
		command = HelpCommand
	} else {
		var source []byte
		var err error
		if IsPiped() {
			source, err = ReadStdin()
			printError(err)
		} else {
			file, err := FindChangelog()
			printError(err)
			source, err = ReadChangelog(file)
			printError(err)
		}
		changelog, err = ParseChangelog(source)
		printError(err)
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
