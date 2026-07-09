// Package feed is the Feed togo plugin: Resource-aggregating social-feed engine: turns N content resources into typed cards with curated round-robin ordering, infinite scroll and masonry.
//
// It self-registers a provider on blank-import and mounts its routes onto the
// kernel. The concrete implementation is ported from the fadymondy.com app under
// internal/server — this scaffold wires the provider + a health route.
package feed

import (
	"net/http"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("feed", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/feed/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"plugin":"feed","status":"ok"}`))
		})
		return nil
	})
}
