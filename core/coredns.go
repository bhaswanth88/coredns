// Package core registers the server and all plugins we support.
package core

import (
	// plug in the server
	_ "github.com/bhaswanth88/coredns/core/dnsserver"
)
