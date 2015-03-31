package main

import (
	"testing"
)

var passingFilenames = []string{"CHANGELOG.yml", "CHANGELOG.yaml", "CHANGE-LOG.yml", "CHANGE_LOG.yml", "changelog.yml", "changelog.yaml"}
var failingFilenames = []string{"CHANGELOG.yml ", " CHANGELOG.yml"}

func TestFilenameRegexp(t *testing.T) {
	for _, filename := range passingFilenames {
		if !REGEXP_FILENAME.MatchString(filename) {
			t.Errorf("Filename %s should be legal", filename)
		}
	}
	for _, filename := range failingFilenames {
		if REGEXP_FILENAME.MatchString(filename) {
			t.Errorf("Filename %s should be illegal", filename)
		}
	}
}
