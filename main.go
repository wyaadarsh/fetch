package main

import (
	"strings"

	"github.com/wyaadarsh/fetch/core"
)

func main() {
	url := "https://cdimage.debian.org/debian-cd/current/amd64/iso-cd/debian-11.7.0-amd64-netinst.iso"
	// extract filename from url
	filename := url[strings.LastIndex(url, "/")+1:]
	info := core.Make_Info(url, filename, 4)
	core.Download(info)
}
