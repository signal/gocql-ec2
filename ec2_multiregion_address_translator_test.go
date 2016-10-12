package gocql_ec2

import (
	"testing"
)

func TestEC2MultiRegionAddressTranslator_NoHostsFoundInReverseLookup(t *testing.T) {
	host, port := newEC2MultiRegionAddressTranslator(NewStaticDNS()).Translate("10.10.220.255", 9042)
	assertEqual(t, "ec2 translated host", "10.10.220.255", host)
	assertEqual(t, "ec2 translated port", 9042, port)
}

func TestEC2MultiRegionAddressTranslator_NoAddrsFoundForReverseLookupHost(t *testing.T) {
	testDNS := NewStaticDNS()
	testDNS.AddAddr("10.10.220.255", "hosty", "pants")
	host, port := newEC2MultiRegionAddressTranslator(testDNS).Translate("10.10.220.255", 9042)
	assertEqual(t, "ec2 translated host", "10.10.220.255", host)
	assertEqual(t, "ec2 translated port", 9042, port)
}

func TestEC2MultiRegionAddressTranslator_TranslateSuccess(t *testing.T) {
	testDNS := NewStaticDNS()
	testDNS.AddAddr("10.10.220.254", "hosty", "pants")
	testDNS.AddHost("hosty", "10.10.220.253", "10.10.220.252")
	testDNS.AddHost("pants", "10.10.220.153", "10.10.220.152")
	host, port := newEC2MultiRegionAddressTranslator(testDNS).Translate("10.10.220.254", 9042)
	assertEqual(t, "ec2 translated host", "10.10.220.253", host)
	assertEqual(t, "ec2 translated port", 9042, port)
}

func TestEC2MultiRegionAddressTranslator_ParseHostIPFails(t *testing.T) {
	testDNS := NewStaticDNS()
	testDNS.AddAddr("10.10.220.254", "hosty", "pants")
	testDNS.AddHost("hosty", "2384908234832") // a bad ip obviously
	testDNS.AddHost("pants", "10.10.220.153", "10.10.220.152")
	host, port := newEC2MultiRegionAddressTranslator(testDNS).Translate("10.10.220.254", 9042)
	assertEqual(t, "ec2 translated host", "10.10.220.254", host)
	assertEqual(t, "ec2 translated port", 9042, port)
}
