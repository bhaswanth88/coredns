package view

import (
	"context"

	"github.com/bhaswanth88/coredns/plugin/metadata"
	"github.com/bhaswanth88/coredns/request"
)

// Metadata implements the metadata.Provider interface.
func (v *View) Metadata(ctx context.Context, state request.Request) context.Context {
	metadata.SetValueFunc(ctx, "view/name", func() string {
		return v.viewName
	})
	return ctx
}
