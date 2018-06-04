package dkmachine

import (
	"encoding/json"
	"fmt"

	"github.com/docker/machine/libmachine/drivers/rpc"
	"github.com/docker/machine/libmachine/host"
)

const bin = "docker-machine"

// Machine ...
type Machine struct {
	CreateOptions *CreateOptions `json:"create_options"`
	HostConfig    *host.Host
}

// Host ...
// TODO: Fix hard cording
func (m *Machine) Host() string {
	u, err := m.HostConfig.Driver.GetURL()
	if err != nil {
		return fmt.Sprintf("dkmachine:error:%s", err.Error())
	}

	return u
}

// CertPath ...
func (m *Machine) CertPath() string {
	return m.HostConfig.HostOptions.AuthOptions.StorePath
}

// Version returns client version. NOT "ConfigVersion".
func (m *Machine) Version() string {
	// return fmt.Sprintf("%d", m.HostConfig.ConfigVersion)
	return ""
}

// GetPrivateIPAddress returns the private IP address of this machine.
// TODO: Hack and Refactor
func (m *Machine) GetPrivateIPAddress() string {

	d, ok := m.HostConfig.Driver.(*rpcdriver.RPCClientDriver)
	if !ok {
		return "/* failed to type cast to RPCClientDriver */"
	}

	b, err := d.GetConfigRaw()
	if err != nil {
		return fmt.Sprintf("/* failed to get raw config of this driver: %v */", err)
	}
	dest := map[string]interface{}{}
	if err := json.Unmarshal(b, &dest); err != nil {
		return fmt.Sprintf("/* failed to decode raw config to map: %v */", err)
	}
	// TODO: Store and reuse this configuration, which is fetched above.

	switch m.HostConfig.DriverName {
	case "amazonec2":
		// FIXME: hard coding... ;(
		return fmt.Sprintf("%v", dest["PrivateIPAddress"])
	case "google":
		// TODO: https://github.com/otiai10/awsub/issues/84
		// m.HostConfig.Driver.GetIP()
		return "/* TODO: GCP instance private IP Address */"
	default:
		return fmt.Sprintf("/* TODO: dkmachine: Private IP Address for %v is not implemented yet */", m.HostConfig.DriverName)
	}
}
