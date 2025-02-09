package test

import (
	"testing"

	"github.com/bhaswanth88/coredns/plugin/test"
)

func TestTempFile(t *testing.T) {
	t.Parallel()
	_, f, e := test.TempFile(".", "test")
	if e != nil {
		t.Fatalf("Failed to create temp file: %s", e)
	}
	defer f()
}
