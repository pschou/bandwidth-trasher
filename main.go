package main

import (
	"os"
	"path"
)

func main() {
	_, exec := path.Split(os.Args[0])
	switch exec[:4] {
	default:
		listener()
	case "send":
		sender()
	case "pull":
		puller()
	}
}
