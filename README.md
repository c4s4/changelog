Changelog
=========

- Project : <https://github.com/c4s4/changelog>.
- Downloads : <https://github.com/c4s4/changelog/releases>.

This is a tool to manage semantic changelogs. A semantic change log is a YAML file where you can extract information regarding releases, such as version, date or summary. Here is a sample semantic changelog:

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

To extract release version, date and summary, you would type:

    $ changelog release version
    1.0.0
    $ changelog release date
    2015-03-30
    $ changelog release summary
    Second release

This tool extracts information from changelog in current directory. You can parse another file using *<* character on command line:

    $ changelog release version < path/to/another/changelog

To get help about this tool, just type:

    $ changelog
    Manage semantic changelog
    
      changelog                      Print this help screen
      changelog release              Check for release
      changelog release date         Print release date
      changelog release date check   Check that release date is correct
      changelog release version      Print release version
      changelog release summary      Print release summary
      changelog to html              Transform changelog to html
      changelog to html stylesheet   Transform to html with a stylesheet
                                     ('style' uses a default stylesheet)
      changelog to markdown          Transform changelog to markdown
    
    The changelog file is searched in current directory. To use a
    different changelog, use < character with its path:
    
      changelog release < path/to/changelog.yml
    
    will check for release a changelog in 'path/to' directory.

Installation
------------

To install *changelog*, you can :

### Install binary

Download latest binary archive at <https://github.com/c4s4/changelog/releases>. Unzip the archive, put the binary of your platform somewhere in your *PATH* and rename it *changelog*.

### Build from sources

Get the project master and build the binary :

        git clone git@github.com:c4s4/changelog
        cd changelog
        go build

This will generate a *changelog* binary for your platform. Put this binary somewhere in your *PATH*.

Changelog Format
----------------

The changelog file is made of a series of entries, one for each release. Each entry must include a release *version* and a release *date*. It may optionally contain a *summary*. For instance:

    - version: 1.2.3
      date:    2015-05-28
      summary: Debug release

After this mandatory header, you may add lists for *added*, *changed*, *deprecated*, *removed*, *fixed* and *security*:

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

This changelog uses [YAML file format](http://yaml.org/spec/1.2/spec.html). Most frequent errors are the following:

- You can't indent with tab characters, this is a syntax error! You *must* use spaces.
- A colon is the character to separate name from value in a map. Thus, if you have a colon in a text, you should surround it with quotes.

Release features
----------------

These features are useful while releasing software: you can extract all release information (such as version, date and summary) from the changelog. You don't have to duplicate release version in changelog and in makefile for instance. You can also check that release version and date formats are correct. Finally you can ensure that release date in changelog is today, thus avoiding a wrong release date in a changelog.

- `changelog release` checks the release entry, which means that changelog must have at least one release block with:
    - A *version* entry is set with *x.y.z* format, where *y* and *z* are optional.
    - A *date* entry is set with ISO format, that is *yyyy-mm-dd*.
    - A *summary* entry is not mandatory.
- `changelog release date` extracts, checks and prints the date of the release.
- `changelog release date check` checks that the release date is current date.
- `changelog release version` extracts, checks and prints release version.
- `changelog release summary` extracts and prints the release summary.

Transformation features
-----------------------

You can transform the YAML changelog into HTML.

- `changelog to html` transforms changelog to HTML and prints it on the console. No stylesheet is applied.
- `changelog to html stylesheet` transforms the changelog to HTML and applies the stylesheet which path is *stylesheet*. The special value *style* for stylesheet applies a default stylesheet.

Usage
-----

You will find an example script that calls *changelog* to perform a release in *sh* directory of the archive.

*Enjoy!*
