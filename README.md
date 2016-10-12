# gocql-ec2

The `gocql_ec2` package provides Amazon EC2 related functionality for the
[gocql cassandra driver](https://github.com/gocql/gocql). Currently this is limited to providing address translation
support for multi-region cassandra ring setups, but who knows what else might be useful. Plus, `gocql_ec2` is a much
more succinct package name than `gocql_ec2_address_translator`.

## Multi-Region Address Translator

The `EC2MultiRegionAddressTranslator` will attempt to resolve any peer's broadcast address to its private-ip.
This works by doing a reverse lookup on the broadcast address -- which should be its public-ip -- then doing an address
lookup for any hostname resulting from the reverse lookup. Within the same ec2 region, this will translate the public-ip
into a private-ip; outside of the peer's ec2 region, only the public-ip will be returned.

This is typically done to save on costs. It's also useful when your "ring masters" disable public-ip access for
connections coming from inside the same network (which they are doing to force you to save on costs).

#### Usage

Configuring `gocql` to use this translator is straight-forward; just configure the `ClusterConfig.AddressTranslator`:

```go
cluster := gocql.NewCluster("node1", "node2")
cluster.AddressTranslator = gocql_ec2.EC2MultiRegionAddressTranslator()
// ...
session := cluster.NewSession()
```

## Testing

Unit tests:

```sh
go test github.com/signal/gocql_ec2/...
```

Integration tests:

```sh
./integration.sh ${GOPATH}
```
