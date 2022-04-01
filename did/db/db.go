package db

import (
	"context"
	"database/sql"
	"log"
)

type Store struct {
	log *log.Logger
	db  sql.DB
}

func NewStore(log *log.Logger, db sql.DB) Store {
	return Store{
		log: log,
		db:  db,
	}
}

// adds a DID to the db
func (s Store) Create(ctx context.Context, did DID) error {
	return nil
}
