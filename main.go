package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	_, exec := path.Split(os.Args[0])
	fmt.Println("called", exec[:4])
	switch exec[:4] {
	case "list", "serv":
		listener()
	case "send":
		sender()
	case "pull":
		puller()
	}
}
