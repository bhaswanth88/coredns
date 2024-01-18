package main

//go:generate go run directives_generate.go
//go:generate go run owners_generate.go

import (
	_ "github.com/bhaswanth88/coredns/core/plugin" // Plug in CoreDNS.
	"github.com/bhaswanth88/coredns/coremain"
)

func main() {
	coremain.Run()
}
