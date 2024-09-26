//go:build !js

package controlhttp

import (
	"context"
	"errors"
)

// Dial connects to the HTTP server at this Dialer's Host:HTTPPort, requests to
// switch to the Tailscale control protocol, and returns an established control
// protocol connection.
//
// If Dial fails to connect using HTTP, it also tries to tunnel over TLS to the
// Dialer's Host:HTTPSPort as a compatibility fallback.
//
// The provided ctx is only used for the initial connection, until
// Dial returns. It does not affect the connection once established.
func (a *Dialer) Dial(ctx context.Context) (*ClientConn, error) {
	if a.Hostname == "" {
		return nil, errors.New("required Dialer.Hostname empty")
	}
	return a.dial(ctx)
}
