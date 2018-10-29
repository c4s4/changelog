package main

import (
	"fmt"
	lib "github.com/c4s4/changelog/lib"
	"os"
	"strconv"
	"strings"
)

func printError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	var changelog lib.Changelog
	var command string
	var args []string
	if len(os.Args) < 2 {
		command = lib.HelpCommand
	} else {
		var source []byte
		var err error
		if lib.IsPiped() {
			source, err = lib.ReadStdin()
			printError(err)
		} else {
			file, err := lib.FindChangelog()
			printError(err)
			source, err = lib.ReadChangelog(file)
			printError(err)
		}
		changelog, err = lib.ParseChangelog(source)
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
	function := lib.CommandMapping[command]
	if function != nil {
		function(changelog, args)
	} else {
		fmt.Printf("Command '%s' unknown\n", command)
		os.Exit(1)
	}
}
