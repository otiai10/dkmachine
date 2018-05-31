package dkmachine

import (
	"fmt"
)

// Create ...
func Create(opt *CreateOptions) (*Machine, error) {

	c, err := NewClient(opt)
	if err != nil {
		return nil, err
	}

	exists, err := c.API.Exists(c.Host.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("a machine with name %s already exists", c.Host.Name)
	}

	driverflags := c.CreateFlags()
	if err := c.Host.Driver.SetConfigFromFlags(driverflags); err != nil {
		return nil, err
	}

	if err := c.API.Create(c.Host); err != nil {
		return nil, err
	}

	if err := c.API.Save(c.Host); err != nil {
		return nil, err
	}

	return &Machine{opt, c.Host}, nil
}
