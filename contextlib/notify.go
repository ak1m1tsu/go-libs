package contextlib

import (
	"context"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

// NotifyContext returns context that handle os signals. If no signals provided
// handle only os.Interrupt signal.
func NotifyContext(parent context.Context, signals ...os.Signal) context.Context {
	if len(signals) == 0 {
		signals = []os.Signal{os.Interrupt}
	}

	ctx, _ := signal.NotifyContext(parent, signals...)
	return ctx
}

// NotifyErrGroup creates a error group with context that handle os signals. If no signals provided
// handle only os.Interrupt signal.
func NotifyErrGroup(parent context.Context, signals ...os.Signal) (*errgroup.Group, context.Context) {
	return errgroup.WithContext(NotifyContext(parent, signals...))
}
