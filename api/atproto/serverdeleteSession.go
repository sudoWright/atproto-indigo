// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.server.deleteSession

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// ServerDeleteSession calls the XRPC method "com.atproto.server.deleteSession".
func ServerDeleteSession(ctx context.Context, c *xrpc.Client) error {
	c.Mux.RLock()
	defer c.Mux.RUnlock()
	if err := c.Do(ctx, xrpc.Procedure, "", "com.atproto.server.deleteSession", nil, nil, nil); err != nil {
		return err
	}

	return nil
}
