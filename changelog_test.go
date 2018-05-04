package main

import (
	"testing"
)

func TestFilenameRegexp(t *testing.T) {
	var passingFilenames = []string{"CHANGELOG.yml", "CHANGELOG.yaml",
		"CHANGE-LOG.yml", "CHANGE_LOG.yml", "changelog.yml", "changelog.yaml"}
	for _, filename := range passingFilenames {
		if !RegexpFilename.MatchString(filename) {
			t.Errorf("Filename %s should be valid", filename)
		}
	}
	var failingFilenames = []string{"CHANGELOG.yml ", " CHANGELOG.yml"}
	for _, filename := range failingFilenames {
		if RegexpFilename.MatchString(filename) {
			t.Errorf("Filename %s should not be valid", filename)
		}
	}
}

func TestDateRegexp(t *testing.T) {
	var passingDates = []string{"2015-04-01"}
	for _, date := range passingDates {
		if !RegexpDate.MatchString(date) {
			t.Errorf("Date %s should be valid", date)
		}
	}
	var failingDates = []string{" 2015-04-01", "2015-04-01 ", "2015/04/01"}
	for _, date := range failingDates {
		if RegexpDate.MatchString(date) {
			t.Errorf("Date %s should not be valid", date)
		}
	}
}

func TestVersionRegexp(t *testing.T) {
	var passingVersions = []string{"1.2.3", "1", "1.2", "100.200.300"}
	for _, version := range passingVersions {
		if !RegexpVersion.MatchString(version) {
			t.Errorf("Version %s should be valid", version)
		}
	}
	var failingVersions = []string{" 1.2.3", "1.2.3 ", ""}
	for _, version := range failingVersions {
		if RegexpVersion.MatchString(version) {
			t.Errorf("Version %s should not be valid", version)
		}
	}
}
