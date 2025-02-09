//go:build gofuzz

package cache

import (
	"github.com/bhaswanth88/coredns/plugin/pkg/fuzz"
)

// Fuzz fuzzes cache.
func Fuzz(data []byte) int {
	return fuzz.Do(New(), data)
}
