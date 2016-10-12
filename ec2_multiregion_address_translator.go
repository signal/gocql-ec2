package gocql_ec2

import (
	"github.com/gocql/gocql"
	"net"
)

// EC2MultiRegionAddressTranslator will do a reverse DNS lookup on the address provided,
// then do a DNS lookup on the first hostname found in order to potentially get the
// private IP of the given address. If no results are found in either lookup,
// the given address is returned. Port is always returned unchanged. Build with the
// gocql_debug tag to see the address translation at work.
func EC2MultiRegionAddressTranslator() gocql.AddressTranslator {
	return newEC2MultiRegionAddressTranslator(builtinDNS{})
}

func newEC2MultiRegionAddressTranslator(dns DNS) gocql.AddressTranslator {
	return gocql.AddressTranslatorFunc(func(addr string, port int) (string, int) {
		names, err := dns.LookupAddr(addr)

		if err != nil || len(names) < 1 {
			return addr, port
		}

		addresses, err := dns.LookupHost(names[0])
		if err != nil || len(addresses) < 1 {
			return addr, port
		}

		newAddr := net.ParseIP(addresses[0])
		if newAddr == nil {
			return addr, port
		}

		return newAddr.String(), port
	})
}
