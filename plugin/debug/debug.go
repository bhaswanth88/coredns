package debug

import (
	"github.com/bhaswanth88/coredns/core/dnsserver"
	"github.com/bhaswanth88/coredns/plugin"
	"github.com/coredns/caddy"
)

func init() { plugin.Register("debug", setup) }

func setup(c *caddy.Controller) error {
	config := dnsserver.GetConfig(c)

	for c.Next() {
		if c.NextArg() {
			return plugin.Error("debug", c.ArgErr())
		}
		config.Debug = true
	}

	return nil
}
