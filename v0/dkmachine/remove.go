package dkmachine

import "os/exec"

// Remove ...
func (m *Machine) Remove() error {
	cmd := exec.Command(bin, "rm", m.Name, "-f", "-y")
	return cmd.Run()
}
