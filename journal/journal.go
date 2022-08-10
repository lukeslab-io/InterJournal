package journal

import (
	"github.com/akrylysov/pogreb"
	"log"
)

type Journal struct {
	DB             *pogreb.DB
	Config         *Config
	ApplicationDir string
	// Keys
	HeadKeys []string
}

// TODO: Implement
type DB interface {
	Open()
	Close()
}

func NewJournal(config *Config) *Journal {
	//TODO: Move this into the a connection function
	db, err := pogreb.Open(config.DBPath, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &Journal{
		DB:       db,
		Config:   config,
		HeadKeys: []string{},
	}
}
