Discord Date Tool
=================

A tool for generating a magic tag that discord will display as a timestamp in
the viewer's timezone.

The program run with no flags will output the current timestamp in the "t" format
or "11:28 AM" (although the AM/PM vs. 24 hour format depends on the user's locale)

This can be altered by adding flags, `-f` for changing the format and adding the
date to the command line in the format `YYYY-MM-DD hh:mm:ss`.

Alternatively the program can be run in GUI mode with the `-g` flag; or if the
name of the executable ends with -gui.

Usage
-----
```
Usage: go-discorddate.exe [-g] [-f format] <date>

  -f string
        Rendering format (default "t")
  -g    Start GUI

Available formats:
         F Saturday, November 4, 2023 11:28:27 AM
         f 4 November 2023 11:28
         R in 2 hours
         D November 4, 2023
         d 04/11/2023
         T 11:28:27 AM
         t 11:28 AM
```

Screenshot
----------

![Screenshot](.assets/screenshot.jpg)
