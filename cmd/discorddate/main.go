package main

import (
	"fmt"
	"time"
	"flag"
	"os"

	"git.shangtai.net/staffan/go-discorddate/internal/dateformat"
	"git.shangtai.net/staffan/go-discorddate/internal/tz"
)

const DATETIME="2006-01-02 15:04:05"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-u] [-f format] <date>\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nAvailable formats:")
		for _, format := range dateformat.Formats {
			fmt.Fprintln(os.Stderr, "\t", format.Key, format.Name)
		}
	}

	var formatkey = flag.String("f", "t", "Rendering format")
	var utc = flag.Bool("u", false, "Input time is in UTC")

	flag.Parse()

	var date time.Time

	if flag.NArg() < 1 {
		date = time.Now()
	} else {
		parsed, err := time.Parse(DATETIME, flag.Arg(0))

		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		if !*utc {
			parsed = tz.AdjustToLocal(parsed)
		}

		date = parsed

	}

	format, err := dateformat.ByKey(*formatkey)

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println(format.Render(date))
}
