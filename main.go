package main

import (
	"flag"
	"github.com/lukemilby/InterJournal/journal"
	"log"
)

// TODO:
// configuration file
// link to data stores
// formatting <Time> <sperateor> <message>
// special highlighting
// will print out tail of your last entries
// sentament anlysis on entry
// status bar
// search operation on entries - melliesearch as an option
// tagging
// auto tagging
// embeded storage
//

func main() {
	// Where are we going to store the file
	// 5 latest entries

	cfg := journal.NewConfig()
	j := journal.NewJournal(cfg)
	/// from the new Journal
	// we load config and create journal
	/// --------- --------- \\\
	// journal takes entries
	// journal middleware for plugins
	// journal to db once the entry has been processed

	// FLAG SETUP HERE
	setup := flag.Bool("setup", false, "setup data and config")

	// Here we parse flags
	flag.Parse()
	message := flag.Args()

	// Create missing directories and files
	if *setup {
		err := j.Setup()
		if err != nil {
			log.Fatal(err)
		}
	}

	// The journal needs to ingest all messages
	j.Run(message)
}
