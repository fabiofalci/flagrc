package main

import (
	"flag"
	"github.com/fabiofalci/flagrc"
)

func main() {
	flagrc.ProcessFlagrc("examplerc")

	providedUsername := flag.String("username", "", "App username")
	providedDate := flag.String("date", "", "App date")

	flag.Parse()

	println("Username: " + *providedUsername)
	println("Date: " + *providedDate)
}
