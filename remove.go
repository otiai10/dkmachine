package dkmachine

import "fmt"

// Remove ...
func (m *Machine) Remove() error {

	c, err := NewClient(m.CreateOptions)
	if err != nil {
		return err
	}

	name := c.Host.Name
	exists, err := c.API.Exists(name)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("a machine with name %s doesn't exist", name)
	}

	h, err := c.API.Load(name)
	if err != nil {
		return fmt.Errorf("failed to load machine configs: %v", err)
	}

	if err := h.Driver.Remove(); err != nil {
		return fmt.Errorf("failed to remove remote machine: %v", err)
	}

	if err := c.API.Remove(name); err != nil {
		return fmt.Errorf("failed to remove local machine config: %v", err)
	}

	return nil

}
