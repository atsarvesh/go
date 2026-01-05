package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {

	part := make([]string, len(ip)) // convert each byte to string

	for i, val := range ip {

		part[i] = strconv.Itoa(int(val))
	}

	return strings.Join(part,".")
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