package shutdown

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func NewShutdown() (*errgroup.Group, context.Context) {
	g, ctx := errgroup.WithContext(context.Background())

	return g, ctx
}
