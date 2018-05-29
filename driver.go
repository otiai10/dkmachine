package dkmachine

import (
	"github.com/docker/machine/drivers/amazonec2"
	"github.com/docker/machine/drivers/google"
	"github.com/docker/machine/libmachine/drivers"
)

func getDriver(opt *CreateOptions) drivers.Driver {
	switch opt.Driver {
	case "amazonec2":
		return driverAmazonEC2(opt)
	case "google":
		return driverGoogleCloud(opt)
	}
	return nil
}

func driverAmazonEC2(opt *CreateOptions) *amazonec2.Driver {
	d := amazonec2.NewDriver("", "")
	d.Region = opt.AmazonEC2Region

	// common
	d.MachineName = opt.Name

	return d
}

func driverGoogleCloud(opt *CreateOptions) *google.Driver {
	d := google.NewDriver("", "")
	d.Zone = opt.GoogleZone
	d.Scopes = opt.GoogleScopes
	d.Project = opt.GoogleProject
	d.DiskSize = opt.GoogleDiskSize

	// common
	d.MachineName = opt.Name

	return d
}
