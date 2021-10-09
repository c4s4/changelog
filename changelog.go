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
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
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
				printError(fmt.Errorf("bad shift '%s'", command))
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
		if err := function(changelog, args); err != nil {
			printError(fmt.Errorf("running command: %v", err))
		}
	} else {
		printError(fmt.Errorf("command '%s' unknown", command))
	}
}
