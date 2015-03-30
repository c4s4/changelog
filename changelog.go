package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
)

const (
    FILENAME = "CHANGELOG.yml"
)

type Release struct {
    Version string
    Date string
    Description string
    Added []string
    Removed []string
    Fixed []string
    Security []string
}

type ChangeLog struct {
    Entries []Release
}

func main() {
    var changelog ChangeLog
    source, err := ioutil.ReadFile(FILENAME)
    if err != nil {
        panic(err)
    }
    err = yaml.Unmarshal(source, &changelog)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Value: %#v\n", changelog)
}