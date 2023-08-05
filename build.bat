windres go-discorddate.rc -o winmanifest_windows_386.syso -O coff
windres go-discorddate.rc -o winmanifest_windows_amd64.syso -O coff
go build 
go build -ldflags -H=windowsgui -o go-discorddate-gui.exe
