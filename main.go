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
		fmt.Printf("Usage: %s [-f format] <date>\n\n", os.Args[0])
		flag.PrintDefaults()
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
