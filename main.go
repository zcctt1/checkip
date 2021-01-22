package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jreisinger/checkip/check"
)

// Version is the default version of checkip.
var Version = "dev"

func main() {
	log.SetPrefix(os.Args[0] + ": ") // prefix program name
	log.SetFlags(0)                  // no timestamp

	flags, err := ParseFlags()
	if err != nil {
		log.Fatalf("%v", err)
	}

	if flags.Version {
		fmt.Println(Version)
		os.Exit(0)
	}

	checks := check.GetAvailable()

	if len(flags.ChecksToRun) > 0 {
		checks = flags.ChecksToRun
	}

	chn := make(chan string)
	for _, chk := range checks {
		go check.Run(chk, flags.IPaddr, chn)
	}
	for range checks {
		fmt.Print(<-chn)
	}
}
