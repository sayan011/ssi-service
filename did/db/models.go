package db

import "time"

// Represents structure for moving data between app and db
type DID struct {
	DID         string
	DateCreated time.Time
	DateUpdated time.Time
}
