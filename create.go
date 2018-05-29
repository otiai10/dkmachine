package dkmachine

import (
	"fmt"
)

// Create ...
func Create(opt *CreateOptions) (*Machine, error) {

	fmt.Println("0001")
	c, err := NewClient(opt)
	if err != nil {
		return nil, err
	}

	fmt.Println("0002")
	exists, err := c.API.Exists(c.Host.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("a machine with name %s already exists", c.Host.Name)
	}

	fmt.Println("0003")
	driverflags := c.CreateFlags()
	if err := c.Host.Driver.SetConfigFromFlags(driverflags); err != nil {
		return nil, err
	}

	fmt.Printf("%v\n", c.Host.Driver.DriverName())
	fmt.Println("0004")
	if err := c.API.Create(c.Host); err != nil {
		return nil, err
	}

	fmt.Println("0005")
	if err := c.API.Save(c.Host); err != nil {
		return nil, err
	}

	return &Machine{opt, c.Host}, nil
}
