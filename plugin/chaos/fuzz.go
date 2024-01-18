//go:build gofuzz

package chaos

import (
	"github.com/bhaswanth88/coredns/plugin/pkg/fuzz"
)

// Fuzz fuzzes cache.
func Fuzz(data []byte) int {
	c := Chaos{}
	return fuzz.Do(c, data)
}
