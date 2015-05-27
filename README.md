Change Log
==========

This is a tool to manage semantic change logs. A semantic change log is a YAML file where you can extract information regarding releases, such as version, date, summary and so on. Here is a sample semantic change log:

    - version: 1.0.0
      date:    2015-03-30
      summary: Second release
      added:
      - First added element.
      - Second added element.
      fixed:
      - Only one fix.
      security:
      - First of course.
      foo:
      - Bar.
    
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

This tool extracts information from a file named *CHANGELOG.yml*, *CHANGELOG.yaml*, *changelog.yml*, *changelog.yaml*, or whatever file which name starts with *change*, ends with *log* and has the extension *.yml* or *.yaml*, in current directory. You can parse a different file using *<* character on command line:

    $ changelog release version < path/to/changelog

To get help about this tool, just type:

    $ changelog
    Manage semantic changelog
    
      changelog                      Print this help screen
      changelog release              Check for release
      changelog release date         Print release date
      changelog release date check   Check that release date wright
      changelog release version      Print release version
      changelog release summary      Print release summary
      changelog to html              Transform changelog to html
      changelog to html stylesheet   Transform tp html with a stylesheet
                                     ('style' uses a default stylesheet)
    
    The changelog file is searched in current directory. To use a
    different changelog, use < character with its path:
    
      changelog release < path/to/changelog.yml
    
    will check for release a changelog in 'path/to' directory.

