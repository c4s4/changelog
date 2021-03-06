# Changelog

[![Build Status](https://travis-ci.org/c4s4/changelog.svg?branch=master)](https://travis-ci.org/c4s4/changelog)
[![Code Quality](https://goreportcard.com/badge/github.com/c4s4/changelog)](https://goreportcard.com/report/github.com/c4s4/changelog)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
<!--
[![Coverage Report](https://coveralls.io/repos/github/c4s4/changelog/badge.svg?branch=master)](https://coveralls.io/github/c4s4/changelog?branch=master)
-->

- Project : <https://github.com/c4s4/changelog>.
- Downloads : <https://github.com/c4s4/changelog/releases>.

This is a tool to manage semantic changelogs. A semantic change log is a YAML file where you can extract information regarding releases, such as version, date or summary. Here is a sample semantic changelog:

```yaml
- version: 1.0.0
  date:    2015-03-30
  summary: Second release
  added:
  - Added element.
  fixed:
  - Fixed element.

- version: 0.1.0
  date:    2015-03-29
  summary: First release
```

To extract release version, date and summary, you would type:

```bash
$ changelog release version
1.0.0
$ changelog release date
2015-03-30
$ changelog release summary
Second release
```

This tool extracts information from changelog in current directory. You can parse another file using *<* character on command line:

```bash
$ changelog release version < path/to/another/changelog
```

To get help about this tool, just type:

```bash
$ changelog
Manage semantic changelog

  changelog                      Print this Help screen
  changelog release              Check for release
  changelog release date         Print release date
  changelog release date check   Check that release date wright
  changelog release version      Print release version
  changelog release summary      Print release summary
  changelog release to markdown  Print release changelog in markdown
  changelog to html              Transform changelog to html
  changelog to html stylesheet   Transform to html with a stylesheet
                                 ('style' uses a default stylesheet)
  changelog to markdown          Transform changelog to markdown

You can add 'next' after changelog command to consider next to last release
instead of the last, or '-N' go go back in past Nth release.

The changelog file is searched in current directory. To use a different
changelog, use < character with its path:

  changelog release < path/to/changelog.yml

will check for release a changelog in 'path/to' directory.
```

## Installation

### Unix users (Linux, BSDs and MacOSX)

Unix users may download and install latest *changelog* release with command:

```bash
$ sh -c "$(curl https://sweetohm.net/dist/changelog/install)"
```

If *curl* is not installed on you system, you might run:

```bash
$ sh -c "$(wget -O - https://sweetohm.net/dist/changelog/install)"
```

**Note:** Some directories are protected, even as *root*, on **MacOSX** (since *El Capitan* release), thus you can't install *changelog* in */usr/bin* for instance.

### Binary package

Otherwise, you can download latest binary archive at <https://github.com/c4s4/dotrun/releases>. Unzip the archive, put the binary of your platform somewhere in your *PATH* and rename it *dotrun*.

### Build from sources

To build *changelog* from source, you will need latest Go version and install [GoYAML](http://gopkg.in/yaml.v2), with following command:

```bash
$ go get -u gopkg.in/yaml.v2
```

Get the project master and build the binary :

```bash
$ git clone git@github.com:c4s4/changelog
$ cd changelog
$ go build
```

This will generate a *changelog* binary for your platform. Put this binary somewhere in your *PATH*.

## Changelog Format

The changelog file is made of a series of entries, one for each release. Each entry must include a release *version* and a release *date*. It may optionally contain a *summary*. For instance:

```yaml
- version: 1.2.3
  date:    2015-05-28
  summary: Debug release
```

After this mandatory header, you may add lists for *added*, *changed*, *deprecated*, *removed*, *fixed* and *security*:

```yaml
- version: 1.2.3
  date:    2015-05-28
  summary: Debug release
  added:
  - First added feature.
  - Second added feature.
  changed:
  - Changed behavior.
  deprecated:
  - Don't use that anymore.
  removed:
  - This is gone.
  fixed:
  - That was fixed.
  security:
  - This security issue was fixed.
```

This changelog uses [YAML file format](http://yaml.org/spec/1.2/spec.html). Most frequent errors are the following:

- You can't indent with tab characters, this is a syntax error! You *must* use spaces.
- A colon is the character to separate name from value in a map. Thus, if you have a colon in a text, you should surround it with quotes.

## Release features

These features are useful while releasing software: you can extract all release information (such as version, date and summary) from the changelog. You don't have to duplicate release version in changelog and in makefile for instance. You can also check that release version and date formats are correct. Finally you can ensure that release date in changelog is today, thus avoiding a wrong release date in a changelog.

- `changelog release` checks the release entry, which means that changelog must have at least one release block with:
    - A *version* entry is set with *x.y.z* format, where *y* and *z* are optional. A *-SNAPSHOT* is also possible at the end of the version number.
    - A *date* entry is set with ISO format, that is *yyyy-mm-dd*.
    - A *summary* entry is not mandatory.
- `changelog release date` extracts, checks and prints the date of the release.
- `changelog release date check` checks that the release date is current date.
- `changelog release version` extracts, checks and prints release version.
- `changelog release summary` extracts and prints the release summary.

## Transformation features

You can transform the YAML changelog into HTML.

- `changelog to html` transforms changelog to HTML and prints it on the console. No stylesheet is applied.
- `changelog to html stylesheet` transforms the changelog to HTML and applies the stylesheet which path is *stylesheet*. The special value *style* for stylesheet applies a default stylesheet.

## Usage

You will find an example script that calls *changelog* to perform a release in *sh* directory of the archive.

*Enjoy!*
