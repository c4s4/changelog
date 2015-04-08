package main

import (
	"fmt"
	"regexp"
	"time"
)

var REGEXP_DATE = regexp.MustCompile(`^\d\d\d\d-\d\d-\d\d$`)
var REGEXP_VERSION = regexp.MustCompile(`^\d+(\.\d+)?(\.\d+)?$`)

func checkRelease(release Release) {
	if release.Version == "" {
		Error(ERROR_RELEASE, "Release version is empty")
	}
	if !REGEXP_VERSION.MatchString(release.Version) {
		Errorf(ERROR_RELEASE, "Release version '%s' is not a valid semantic version number",
			release.Version)
	}
	if release.Date == "" {
		Error(ERROR_RELEASE, "Release date is empty")
	}
	if !REGEXP_DATE.MatchString(release.Date) {
		Errorf(ERROR_RELEASE, "Release date '%s' is not valid ISO format", release.Date)
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
			if len(args) > 1 {
				date := time.Now().Local().Format("2006-01-02")
				if date != (*changelog)[0].Date {
					Errorf(ERROR_RELEASE, "Release date %s is wrong (should be %s)",
						(*changelog)[0].Date, date)
				}
			} else {
				fmt.Println((*changelog)[0].Date)
			}
		} else if args[0] == "version" {
			fmt.Println((*changelog)[0].Version)
		} else {
			Errorf(ERROR_RELEASE, "Unknown release argument %s", args[0])
		}
	}
}
