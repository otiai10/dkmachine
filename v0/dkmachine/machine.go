package dkmachine

import (
	"fmt"
	"net/url"
)

const bin = "docker-machine"

// Machine ...
type Machine struct {
	*Inspection   `json:",inline"`
	*Env          `json:",inline"`
	CreateOptions *CreateOptions `json:"create_options"`
}

// Host ...
// TODO: Fix hard cording
func (m *Machine) Host() string {

	if m.Inspection == nil {
		return "dkmachine:unknown_host"
	}

	u, err := url.Parse(m.Env.Host)
	if err != nil {
		return fmt.Sprintf("dkmachine:error:%s", err.Error())
	}
	if u.Port() == "" {
		u.Host += ":2376"
	}

	return u.String()
}

// CertPath ...
func (m *Machine) CertPath() string {
	return m.Env.CertPath
}

// Version ...
func (m *Machine) Version() string {
	return ""
}
