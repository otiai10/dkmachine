package dkmachine

import (
	"fmt"

	"github.com/docker/machine/libmachine/host"
)

const bin = "docker-machine"

// Machine ...
type Machine struct {
	CreateOptions *CreateOptions `json:"create_options"`
	host          *host.Host
}

// Host ...
// TODO: Fix hard cording
func (m *Machine) Host() string {

	u, err := m.host.Driver.GetURL()
	if err != nil {
		return fmt.Sprintf("dkmachine:error:%s", err.Error())
	}

	return u
}

// CertPath ...
func (m *Machine) CertPath() string {
	return m.host.AuthOptions().CaCertPath
}

// Version ...
func (m *Machine) Version() string {
	return fmt.Sprintf("%d", m.host.ConfigVersion)
}
