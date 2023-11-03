Discord Date Tool
=================

A tool for generating a magic tag that discord will display as a timestamp in
the viewer's timezone.

The program run with no flags will output the current timestamp in the "t"
format or "11:28 AM" (although the AM/PM vs. 24 hour format depends on the
user's locale)

This can be altered by adding flags, `-f` for changing the format and adding
the date to the command line in the format `YYYY-MM-DD hh:mm:ss`.

Usage
-----
```
Usage: discorddate [-f format] <date>

  -f string
        Rendering format (default "t")

Available formats:
         F Saturday, November 4, 2023 11:28:27 AM
         f 4 November 2023 11:28
         R in 2 hours
         D November 4, 2023
         d 04/11/2023
         T 11:28:27 AM
         t 11:28 AM
```

GUI
---

There is a gui version using [golang-ui](https://github.com/libui-ng/golang-ui)
 in the `cmd/discorddate-gui` directory.

On windows, an extra step of `go generate ./cmd/dockerdate-gui`
(or `go generate ./...` if you like) is required to build the windows binary
resource files which gives the binary an icon and some other random info like
author and version.

Building
--------

To build all binaries, run
```
% go generate ./...
% go build -o . ./...
```

Build only console binary
```
% go build ./cmd/dockerdate
```

Build only GUI binary
```
% go generate ./...
% go build ./cmd/dockerdate-gui
```

Screenshot
----------

![Screenshot](.assets/screenshot.jpg)
