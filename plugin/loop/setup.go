package loop

import (
	"net"
	"strconv"
	"time"

	"github.com/bhaswanth88/coredns/core/dnsserver"
	"github.com/bhaswanth88/coredns/plugin"
	"github.com/bhaswanth88/coredns/plugin/pkg/dnsutil"
	"github.com/bhaswanth88/coredns/plugin/pkg/rand"
	"github.com/coredns/caddy"
)

func init() { plugin.Register("loop", setup) }

func setup(c *caddy.Controller) error {
	l, err := parse(c)
	if err != nil {
		return plugin.Error("loop", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		l.Next = next
		return l
	})

	// Send query to ourselves and see if it end up with us again.
	c.OnStartup(func() error {
		// Another Go function, otherwise we block startup and can't send the packet.
		go func() {
			deadline := time.Now().Add(30 * time.Second)
			conf := dnsserver.GetConfig(c)
			lh := conf.ListenHosts[0]
			addr := net.JoinHostPort(lh, conf.Port)

			for time.Now().Before(deadline) {
				l.setAddress(addr)
				if _, err := l.exchange(addr); err != nil {
					l.reset()
					time.Sleep(1 * time.Second)
					continue
				}

				go func() {
					time.Sleep(2 * time.Second)
					l.setDisabled()
				}()

				break
			}
			l.setDisabled()
		}()
		return nil
	})

	return nil
}

func parse(c *caddy.Controller) (*Loop, error) {
	i := 0
	zones := []string{"."}
	for c.Next() {
		if i > 0 {
			return nil, plugin.ErrOnce
		}
		i++
		if c.NextArg() {
			return nil, c.ArgErr()
		}

		if len(c.ServerBlockKeys) > 0 {
			zones = plugin.Host(c.ServerBlockKeys[0]).NormalizeExact()
		}
	}
	return New(zones[0]), nil
}

// qname returns a random name. <rand.Int()>.<rand.Int().<zone>.
func qname(zone string) string {
	l1 := strconv.Itoa(r.Int())
	l2 := strconv.Itoa(r.Int())

	return dnsutil.Join(l1, l2, zone)
}

var r = rand.New(time.Now().UnixNano())
