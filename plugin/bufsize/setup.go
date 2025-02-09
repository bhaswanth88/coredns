package bufsize

import (
	"strconv"

	"github.com/bhaswanth88/coredns/core/dnsserver"
	"github.com/bhaswanth88/coredns/plugin"
	"github.com/coredns/caddy"
)

func init() { plugin.Register("bufsize", setup) }

func setup(c *caddy.Controller) error {
	bufsize, err := parse(c)
	if err != nil {
		return plugin.Error("bufsize", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Bufsize{Next: next, Size: bufsize}
	})

	return nil
}

func parse(c *caddy.Controller) (int, error) {
	// value from http://www.dnsflagday.net/2020/
	const defaultBufSize = 1232
	for c.Next() {
		args := c.RemainingArgs()
		switch len(args) {
		case 0:
			// Nothing specified; use defaultBufSize
			return defaultBufSize, nil
		case 1:
			// Specified value is needed to verify
			bufsize, err := strconv.Atoi(args[0])
			if err != nil {
				return -1, plugin.Error("bufsize", c.ArgErr())
			}
			// Follows RFC 6891
			if bufsize < 512 || bufsize > 4096 {
				return -1, plugin.Error("bufsize", c.ArgErr())
			}
			return bufsize, nil
		default:
			// Only 1 argument is acceptable
			return -1, plugin.Error("bufsize", c.ArgErr())
		}
	}
	return -1, plugin.Error("bufsize", c.ArgErr())
}
