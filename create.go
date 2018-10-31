package dkmachine

import (
	"bytes"
	"fmt"

	"github.com/docker/machine/libmachine/log"
)

// Create ...
func Create(opt *CreateOptions) (*Machine, error) {

	// FIXME: docker/machine/libmachine/log uses global static logger
	b := bytes.NewBuffer(nil)
	log.SetOutWriter(b)
	log.SetErrWriter(b)

	machine := &Machine{CreateOptions: opt}

	c, err := NewClient(opt)
	if err != nil {
		return machine, err
	}
	defer func() {
		machine.HostConfig = c.Host
	}()

	exists, err := c.API.Exists(c.Host.Name)
	if err != nil {
		return machine, err
	}
	if exists {
		return machine, fmt.Errorf("a machine with name %s already exists", c.Host.Name)
	}

	driverflags := c.CreateFlags()
	if err := c.Host.Driver.SetConfigFromFlags(driverflags); err != nil {
		return machine, err
	}

	if err := c.API.Create(c.Host); err != nil {
		return machine, err
	}

	if err := c.API.Save(c.Host); err != nil {
		return machine, err
	}

	return machine, nil
}
