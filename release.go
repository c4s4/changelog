package main

import (
	"fmt"
	"regexp"
	"time"
)

// RegexpDate is a regexp for date
var RegexpDate = regexp.MustCompile(`^\d\d\d\d-\d\d-\d\d$`)

// RegexSuffixes is a regexp for version suffixes
var RegexSuffixes = `SNAPSHOT|ALPHA|BETA|RC|snapshot|alpha|beta|rc`

// RegexpVersion is a regexp for version
var RegexpVersion = regexp.MustCompile(`^\d+(\.\d+)*(-(` + RegexSuffixes + `)(-\d+)?)?$`)

func checkRelease(release Release) error {
	if release.Version == "" {
		return fmt.Errorf("Release version is empty")
	}
	if !RegexpVersion.MatchString(release.Version) {
		return fmt.Errorf("Release version '%s' is not a valid semantic version number",
			release.Version)
	}
	if release.Date == "" {
		return fmt.Errorf("Release date is empty")
	}
	if !RegexpDate.MatchString(release.Date) {
		return fmt.Errorf("Release date '%s' is not valid ISO format", release.Date)
	}
	return nil
}

func checkChangelog(changelog Changelog) error {
	if len(changelog) == 0 {
		return fmt.Errorf("Release is empy")
	}
	for _, release := range changelog {
		err := checkRelease(release)
		if err != nil {
			return err
		}
	}
	return nil
}

func release(changelog Changelog, args []string) error {
	checkChangelog(changelog)
	if len(args) > 0 {
		if args[0] == "summary" {
			fmt.Println((changelog)[0].Summary)
		} else if args[0] == "date" {
			if len(args) > 1 {
				date := time.Now().Local().Format("2006-01-02")
				if date != (changelog)[0].Date {
					return fmt.Errorf("Release date %s is wrong (should be %s)",
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
			return fmt.Errorf("Unknown release argument %s", args[0])
		}
	}
	return nil
}
