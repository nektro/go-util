package main

import (
	"os"

	"github.com/nektro/go-util/mbpp"
)

func main() {
	mbpp.Init(4)

	fileName := "debian-10.4.0-amd64-netinst.iso"
	mbpp.CreateDownloadJob("http://ftp.us.debian.org/debian-cdimage/current/amd64/iso-cd/"+fileName, fileName, nil)

	mbpp.Wait()
	os.Remove(fileName)
}
