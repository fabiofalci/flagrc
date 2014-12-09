Flagrc
------

Flagrc enable arguments to be read from a file, similiar to [.ackrc](http://beyondgrep.com/documentation/).

Usage
-----

Examplerc file

	-username=myusername
	-date=2014-01-12

Go code

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

Output

	$ go run example2.go
	Username: myusername
	Date: 2014-01-12

	$ go run example2.go -date=1970-01-01
	Username: myusername
	Date: 1970-01-01

	$ go run example2.go -username=other
	Username: other
	Date: 2014-01-12

	$ go run example2.go -date=1970-01-01 -username=other
	Username: other
	Date: 1970-01-01
