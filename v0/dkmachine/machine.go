package dkmachine

import "fmt"

const bin = "docker-machine"

// Machine ...
type Machine struct {
	*Inspection   `json:",inline"`
	CreateOptions *CreateOptions `json:"create_options"`
}

// Host ...
// TODO: Fix hard cording
func (m *Machine) Host() string {
	if m.Inspection == nil {
		return "dkmachine:unknown_host"
	}
	return fmt.Sprintf("tcp://%s:%s", m.Inspection.Driver.IPAddress, "2376")
}

// CertPath ...
func (m *Machine) CertPath() string {
	if m.Inspection == nil {
		return "dkmachine:unknown_certpath"
	}
	return m.Inspection.HostOptions.AuthOptions.StorePath
}
