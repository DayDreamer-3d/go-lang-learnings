package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
	var ipString []string

	for _, v := range i {
		intV := strconv.Itoa(int(v))
		ipString = append(ipString, intV)
	}

	return strings.Join(ipString, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
