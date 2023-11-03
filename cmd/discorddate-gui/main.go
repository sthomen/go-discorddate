package main

//go:generate windres go-discorddate.rc -o winmanifest_windows_386.syso -O coff
//go:generate windres go-discorddate.rc -o winmanifest_windows_amd64.syso -O coff

import (
	"time"

	"git.shangtai.net/staffan/go-discorddate/internal/dateformat"
	"git.shangtai.net/staffan/go-discorddate/internal/gui"
)

func main() {
	format, _ := dateformat.ByKey("t")
	gui.MainWindow(time.Now(), format)
}
