// Package deprecated is used when we deprecated plugin. In plugin.cfg just go from
//
// startup:github.com/coredns/caddy/startupshutdown
//
// To:
//
// startup:deprecated
//
// And things should work as expected. This means starting CoreDNS will fail with an error. We can only
// point to the release notes to details what next steps a user should take. I.e. there is no way to add this
// to the error generated.
package deprecated

import (
	"errors"

	"github.com/bhaswanth88/coredns/plugin"
	"github.com/coredns/caddy"
)

// removed has the names of the plugins that need to error on startup.
var removed = []string{""}

func setup(c *caddy.Controller) error {
	c.Next()
	x := c.Val()
	return plugin.Error(x, errors.New("this plugin has been deprecated"))
}

func init() {
	for _, plug := range removed {
		plugin.Register(plug, setup)
	}
}
