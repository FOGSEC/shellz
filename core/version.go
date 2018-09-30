package core

const (
	Name    = "shellz"
	Version = "ù1.0.4"
	Author  = "Simone 'evilsocket' Margaritelli"
	Website = "https://evilsocket.net/"
)

var (
	Banner = Blue(`
  ██████  ██░ ██ ▓█████  ██▓     ██▓    ▒███████▒
▒██    ▒ ▓██░ ██▒▓█   ▀ ▓██▒    ▓██▒    ▒ ▒ ▒ ▄▀░
░ ▓██▄   ▒██▀▀██░▒███   ▒██░    ▒██░    ░ ▒ ▄▀▒░ 
  ▒   ██▒░▓█ ░██ ▒▓█  ▄ ▒██░    ▒██░      ▄▀▒   ░
▒██████▒▒░▓█▒░██▓░▒████▒░██████▒░██████▒▒███████▒
▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░░ ▒░ ░░ ▒░▓  ░░ ▒░▓  ░░▒▒ ▓░▒░▒
░ ░▒  ░ ░ ▒ ░▒░ ░ ░ ░  ░░ ░ ▒  ░░ ░ ▒  ░░░▒ ▒ ░ ▒
░  ░  ░   ░  ░░ ░   ░     ░ ░     ░ ░   ░ ░ ░ ░ ░
      ░   ░  ░  ░   ░  ░    ░  ░    ░  ░  ░ ░    
                                        ░  `) +
		"v" + Version + "\n" +
		Dim("Made with ") + Red("❤") + Dim("  by "+Author) +
		"\n"
)
