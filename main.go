package main

import (
	"fmt"
	"os"

	log "github.com/YoungPioneers/blog4go"
)

func main() {
	println("hello")

	err := log.NewWriterFromConfigAsFile("config.xml")
	if nil != err {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer log.Close()
	log.Warn("what???")
}
