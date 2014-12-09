package main

import (
	"github.com/fabiofalci/flagrc"
	"os"
)

func main() {
	flagrc.ProcessFlagrc("examplerc")

	println(os.Args[1])
	println(os.Args[2])
}
