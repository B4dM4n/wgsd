package main

import (
	"fmt"

	_ "github.com/jwhited/wgsd"

	"github.com/coredns/coredns/core/dnsserver"
	_ "github.com/coredns/coredns/core/plugin"
	"github.com/coredns/coredns/coremain"
)

func find(a []string, s string) int {
	for i, n := range a {
		if n == s {
			return i
		}
	}
	panic(fmt.Sprintf("%q directive not found in %q", s, a))
}

func init() {
	d := dnsserver.Directives
	i := find(d, "template")

	d = append(d, "")
	copy(d[i+1:], d[i:])
	d[i] = "wgsd"

	dnsserver.Directives = d
}

func main() {
	coremain.Run()
}
