//go:build gofuzz

package whoami

import (
	"github.com/bhaswanth88/coredns/plugin/pkg/fuzz"
)

// Fuzz fuzzes cache.
func Fuzz(data []byte) int {
	w := Whoami{}
	return fuzz.Do(w, data)
}
