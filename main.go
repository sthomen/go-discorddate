package main

import (
	"fmt"
	"time"
	"git.shangtai.net/staffan/go-discorddate/dateformat"
)

func main() {
	var format = dateformat.Get(dateformat.FORMAT_t)
	fmt.Printf("Test: %s", format.Render(time.Now()))
}
