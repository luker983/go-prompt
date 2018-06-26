// +build windows

package prompt

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

type SignalHandler struct {
	SigWinch chan struct{}
}

func NewSignalHandler() *SignalHandler {
	return &SignalHandler{}
}

func (sh *SignalHandler) Run(ctx context.Context, cancel context.CancelFunc) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(
		sigchan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	for {
		select {
		case <-ctx.Done():
			return
		case s := <-sigchan:
			cancel()
		}
	}
}
