package dkmachine

import (
	"fmt"
	"net/url"
)

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
	u, err := url.Parse(m.Inspection.Driver.IPAddress)
	if err != nil {
		return fmt.Sprintf("dkmachine:%v", err)
	}
	if u.Scheme == "" {
		u.Scheme = "tcp"
	}
	if u.Port() == "" {
		u.Host += ":2376"
	}
	return u.String()
}

// CertPath ...
func (m *Machine) CertPath() string {
	if m.Inspection == nil {
		return "dkmachine:unknown_certpath"
	}
	return m.Inspection.HostOptions.AuthOptions.StorePath
}
