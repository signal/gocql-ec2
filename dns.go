package gocql_ec2

import "net"

type DNS interface {
	LookupAddr(addr string) (names []string, err error)
	LookupHost(host string) (addrs []string, err error)
}

type builtinDNS struct {}

func (d builtinDNS) LookupAddr(addr string) ([]string, error) {
	return net.LookupAddr(addr)
}

func (d builtinDNS) LookupHost(host string) ([]string, error) {
	return net.LookupHost(host)
}

