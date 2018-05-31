package dkmachine

import (
	"github.com/docker/machine/commands/mcndirs"
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

// Create *amazonec2.Driver.
// See also https://godoc.org/github.com/docker/machine/drivers/amazonec2#Driver.SetConfigFromFlags
func driverAmazonEC2(opt *CreateOptions) *amazonec2.Driver {

	d := amazonec2.NewDriver("", "")

	// common
	d.MachineName = opt.Name
	d.StorePath = mcndirs.GetBaseDir()

	d.Region = opt.AmazonEC2Region
	d.SecurityGroupName = opt.AmazonEC2SecurityGroup
	d.InstanceType = opt.AmazonEC2InstanceType
	d.IamInstanceProfile = opt.AmazonEC2IAMInstanceProfile
	d.RootSize = int64(opt.AmazonEC2RootSize)
	d.RequestSpotInstance = opt.AmazonEC2RequestSpotInstance

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
	d.StorePath = mcndirs.GetBaseDir()

	return d
}
