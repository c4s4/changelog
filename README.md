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

