package main

import (
	"github.com/colm2/impressive"
	"log"
	"os"
)

const (
	usage    = "usage: impressive email@umail.ucc.ie password"
	errormsg = "An error occured generating your calendar. Did you enter your username and password?"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal(usage)
	}
	ical, err := impressive.GetICal(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(errormsg + "\n" + usage)
	}

	log.Print(ical)
}
