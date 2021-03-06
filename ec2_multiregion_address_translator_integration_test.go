// +build integration

package gocql_ec2

import (
	"testing"
	"net"
)

// see integration.sh for what addresses are available in /etc/hosts

func TestIntegration_EC2MultiRegionAddressTranslator_NoHostsFoundInReverseLookup(t *testing.T) {
	host, port := EC2MultiRegionAddressTranslator().Translate(net.ParseIP("10.10.220.155"), 9042)
	assertIPs(t, "ec2 translated host", net.ParseIP("10.10.220.155"), host)
	assertEqual(t, "ec2 translated port", 9042, port)
}

func TestIntegration_EC2MultiRegionAddressTranslator_TranslateSuccess(t *testing.T) {
	host, port := EC2MultiRegionAddressTranslator().Translate(net.ParseIP("10.10.220.154"), 9042)
	assertIPs(t, "ec2 translated host", net.ParseIP("10.10.220.153"), host)
	assertEqual(t, "ec2 translated port", 9042, port)
}
