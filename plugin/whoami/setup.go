package whoami

import (
	"github.com/bhaswanth88/coredns/core/dnsserver"
	"github.com/bhaswanth88/coredns/plugin"
	"github.com/coredns/caddy"
)

func init() { plugin.Register("whoami", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // 'whoami'
	if c.NextArg() {
		return plugin.Error("whoami", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Whoami{}
	})

	return nil
}
