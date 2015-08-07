package execdriver

import "github.com/docker/docker/pkg/nat"

// NetworkInterface contains all network configs for a driver
type NetworkInterface struct {
	CommonNetworkInterface

	// Fields below here are platform specific.

	// PortBindings is the port mapping between the exposed port in the
	// container and the port on the host.
	PortBindings nat.PortMap
}
