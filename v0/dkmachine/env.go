package dkmachine

import (
	"bytes"
	"os/exec"
	"regexp"
)

// Env ...
type Env struct {
	TLSVerify   string
	Host        string
	CertPath    string
	MachineName string
}

// Decode ...
func (e *Env) Decode(b []byte) error {
	exp := regexp.MustCompile(`^export ([A-Z_]+)="(.+)"$`)
	for _, line := range bytes.Split(b, []byte("\n")) {
		match := exp.FindSubmatch(line)
		if len(match) < 3 {
			continue
		}
		switch string(match[1]) {
		case "DOCKER_TLS_VERIFY":
			e.TLSVerify = string(match[2])
		case "DOCKER_HOST":
			e.Host = string(match[2])
		case "DOCKER_CERT_PATH":
			e.CertPath = string(match[2])
		case "DOCKER_MACHINE_NAME":
			e.MachineName = string(match[2])
		}
	}
	return nil
}

// GetEnv ...
func (m *Machine) GetEnv() (*Env, error) {
	cmd := exec.Command(bin, "env", m.Name)
	o, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	env := new(Env)
	if err := env.Decode(o); err != nil {
		return nil, err
	}
	m.Env = env
	return env, nil
}
