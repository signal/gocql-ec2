package gocql_ec2

import (
	"net"
	"testing"
)

func TestEC2MultiRegionAddressTranslator_NoHostsFoundInReverseLookup(t *testing.T) {
	host := net.ParseIP("10.10.220.255")
	actualHost, actualPort := newEC2MultiRegionAddressTranslator(NewStaticDNS()).Translate(host, 9042)
	assertIPs(t, "ec2 translated host", host, actualHost)
	assertEqual(t, "ec2 translated port", 9042, actualPort)
}

func TestEC2MultiRegionAddressTranslator_NoAddrsFoundForReverseLookupHost(t *testing.T) {
	testDNS := NewStaticDNS()
	testDNS.AddAddr("10.10.220.255", "hosty", "pants")
	host := net.ParseIP("10.10.220.255")
	actualHost, actualPort := newEC2MultiRegionAddressTranslator(testDNS).Translate(host, 9042)
	assertIPs(t, "ec2 translated host", host, actualHost)
	assertEqual(t, "ec2 translated port", 9042, actualPort)
}

func TestEC2MultiRegionAddressTranslator_TranslateSuccess(t *testing.T) {
	testDNS := NewStaticDNS()
	testDNS.AddAddr("10.10.220.254", "hosty", "pants")
	testDNS.AddHost("hosty", "10.10.220.253", "10.10.220.252")
	testDNS.AddHost("pants", "10.10.220.153", "10.10.220.152")
	actualHost, actualPort := newEC2MultiRegionAddressTranslator(testDNS).Translate(net.ParseIP("10.10.220.254"), 9042)
	assertIPs(t, "ec2 translated host", net.ParseIP("10.10.220.253"), actualHost)
	assertEqual(t, "ec2 translated port", 9042, actualPort)
}

func TestEC2MultiRegionAddressTranslator_ParseHostIPFails(t *testing.T) {
	testDNS := NewStaticDNS()
	testDNS.AddAddr("10.10.220.254", "hosty", "pants")
	testDNS.AddHost("hosty", "2384908234832") // a bad ip obviously
	testDNS.AddHost("pants", "10.10.220.153", "10.10.220.152")
	actualHost, actualPort := newEC2MultiRegionAddressTranslator(testDNS).Translate(net.ParseIP("10.10.220.254"), 9042)
	assertIPs(t, "ec2 translated host", net.ParseIP("10.10.220.254"), actualHost)
	assertEqual(t, "ec2 translated port", 9042, actualPort)
}
