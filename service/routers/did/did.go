package didrouter

import (
	"context"
	"net/http"

	"github.com/tbd54566975/vc-service/did"
)

type Handlers struct {
	DID did.State
}

// Create adds a new DID .
func (h Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// things
	// h.DID.Create()
	return nil
}
