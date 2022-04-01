package did

import (
	"unsafe"

	"github.com/tbd54566975/vc-service/did/db"
)

// NewDID contains information needed to create a new DID.
type NewDID struct {
	Method string `json:"method"`
}

// this struct represents the application representation of DID
type DID struct {
	DID string
}

func toDID(dbDID db.DID) DID {
	did := (*DID)(unsafe.Pointer(&dbDID))
	return *did
}
