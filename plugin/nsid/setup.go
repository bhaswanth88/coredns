package nsid

import (
	"os"
	"strings"

	"github.com/bhaswanth88/coredns/core/dnsserver"
	"github.com/bhaswanth88/coredns/plugin"
	"github.com/coredns/caddy"
)

func init() { plugin.Register("nsid", setup) }

func setup(c *caddy.Controller) error {
	nsid, err := nsidParse(c)
	if err != nil {
		return plugin.Error("nsid", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Nsid{Next: next, Data: nsid}
	})

	return nil
}

func nsidParse(c *caddy.Controller) (string, error) {
	// Use hostname as the default
	nsid, err := os.Hostname()
	if err != nil {
		nsid = "localhost"
	}
	i := 0
	for c.Next() {
		if i > 0 {
			return nsid, plugin.ErrOnce
		}
		i++
		args := c.RemainingArgs()
		if len(args) > 0 {
			nsid = strings.Join(args, " ")
		}
	}
	return nsid, nil
}
