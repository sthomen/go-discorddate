package main

import (
	"fmt"
	"time"
	"flag"
	"os"
	"git.shangtai.net/staffan/go-discorddate/dateformat"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-f format] <date>\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nAvailable formats:")
		for key, name := range dateformat.Formats() {
			fmt.Fprintln(os.Stderr, "\t", key, name)
		}
	}

	var formatkey = flag.String("f", "t", "Rendering format")

	flag.Parse()

	var date time.Time

	if flag.NArg() < 1 {
		date = time.Now()
	} else {
		parsed, err := time.Parse(time.DateTime, flag.Arg(0))

		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		date = parsed
	}


	format, err := dateformat.Get(*formatkey)

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println(format.Render(date))
}
