package main

import (
	"fmt"
	"regexp"
	"time"
)

// RegexpDate is a regexp for date
var RegexpDate = regexp.MustCompile(`^\d\d\d\d-\d\d-\d\d$`)

// RegexpVersion is a regexp for version
var RegexpVersion = regexp.MustCompile(`^\d+(\.\d+)*(-(SNAPSHOT|ALPHA|BETA|snapshot|alpha|beta)(-\d+)?)?$`)

func checkRelease(release Release) {
	if release.Version == "" {
		Error(ErrorRelease, "Release version is empty")
	}
	if !RegexpVersion.MatchString(release.Version) {
		Errorf(ErrorRelease, "Release version '%s' is not a valid semantic version number",
			release.Version)
	}
	if release.Date == "" {
		Error(ErrorRelease, "Release date is empty")
	}
	if !RegexpDate.MatchString(release.Date) {
		Errorf(ErrorRelease, "Release date '%s' is not valid ISO format", release.Date)
	}
}

func checkChangelog(changelog Changelog) {
	if len(changelog) == 0 {
		Error(ErrorRelease, "Release is empy")
	}
	for _, release := range changelog {
		checkRelease(release)
	}
}

func release(changelog Changelog, args []string) {
	checkChangelog(changelog)
	if len(args) > 0 {
		if args[0] == "summary" {
			fmt.Println((changelog)[0].Summary)
		} else if args[0] == "date" {
			if len(args) > 1 {
				date := time.Now().Local().Format("2006-01-02")
				if date != (changelog)[0].Date {
					Errorf(ErrorRelease, "Release date %s is wrong (should be %s)",
						(changelog)[0].Date, date)
				}
			} else {
				fmt.Println((changelog)[0].Date)
			}
		} else if args[0] == "version" {
			fmt.Println((changelog)[0].Version)
		} else if args[0] == "to" {
			releaseToMarkdown((changelog)[0])
		} else {
			Errorf(ErrorRelease, "Unknown release argument %s", args[0])
		}
	}
}
