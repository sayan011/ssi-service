// this package acts as an API that wraps the data layer
package did

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/tbd54566975/vc-service/did/db"
)

type State struct {
	store db.Store
}

// NewState constructs state for the DID API.
func NewState(log *log.Logger, store sql.DB) State {
	return State{
		store: db.NewStore(log, store),
	}
}

// Create inserts a new user into the database.
func (s State) Create(ctx context.Context, nd NewDID, now time.Time) (string, error) {
	s.Create()
}
