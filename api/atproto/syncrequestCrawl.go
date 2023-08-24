// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.sync.requestCrawl

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// SyncRequestCrawl_Input is the input argument to a com.atproto.sync.requestCrawl call.
type SyncRequestCrawl_Input struct {
	// hostname: Hostname of the service that is requesting to be crawled.
	Hostname string `json:"hostname" cborgen:"hostname"`
}

// SyncRequestCrawl calls the XRPC method "com.atproto.sync.requestCrawl".
func SyncRequestCrawl(ctx context.Context, c *xrpc.Client, input *SyncRequestCrawl_Input) error {
	c.Mux.RLock()
	defer c.Mux.RUnlock()
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.sync.requestCrawl", nil, input, nil); err != nil {
		return err
	}

	return nil
}
