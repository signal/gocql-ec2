package gocql_ec2

import (
	"testing"
)

func assertEqual(t *testing.T, description string, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("expected %s to be (%+v) but was (%+v) instead", description, expected, actual)
	}
}

// a DNS that you can control

type addrsList []string
type addrsMap map[string]addrsList

type staticDNS struct {
	addrToHostsMap addrsMap
	hostToAddrsMap addrsMap
}

func NewStaticDNS() *staticDNS {
	return &staticDNS{
		addrToHostsMap: make(addrsMap),
		hostToAddrsMap: make(addrsMap),
	}
}

func (d *staticDNS) LookupAddr(addr string) ([]string, error) {
	if names, ok := d.addrToHostsMap[addr]; ok {
		return names, nil
	}
	return []string{}, nil
}

func (d *staticDNS) LookupHost(host string) ([]string, error) {
	if addrs, ok := d.hostToAddrsMap[host]; ok {
		return addrs, nil
	}
	return []string{}, nil
}

func (d *staticDNS) AddAddr(addr string, hosts ...string) {
	curHosts, ok := d.addrToHostsMap[addr]
	if !ok {
		curHosts = make(addrsList, 0, len(hosts))
	}
	d.addrToHostsMap[addr] = append(curHosts, hosts...)
}

func (d *staticDNS) AddHost(host string, addrs ...string) {
	curAddrs, ok := d.hostToAddrsMap[host]
	if !ok {
		curAddrs = make(addrsList, 0, len(addrs))
	}
	d.hostToAddrsMap[host] = append(curAddrs, addrs...)
}
