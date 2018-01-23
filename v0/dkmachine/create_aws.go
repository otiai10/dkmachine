package dkmachine

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	// AWSMachineNameExpression ...
	AWSMachineNameExpression = regexp.MustCompile("^[0-9a-zA-Z.-]+$")
)

// ArgsForAmazonEC2 ...
func (opt *CreateOptions) ArgsForAmazonEC2() []string {
	args := []string{
		"--driver", "amazonec2",
	}
	if opt.AmazonEC2InstanceType != "" {
		args = append(args, "--amazonec2-instance-type", opt.AmazonEC2InstanceType)
	}
	if opt.AmazonEC2Region != "" {
		args = append(args, "--amazonec2-region", opt.AmazonEC2Region)
	}
	if opt.AmazonEC2IAMInstanceProfile != "" {
		args = append(args, "--amazonec2-iam-instance-profile", opt.AmazonEC2IAMInstanceProfile)
	}
	if opt.AmazonEC2SecurityGroup != "" {
		args = append(args, "--amazonec2-security-group", opt.AmazonEC2SecurityGroup)
	}
	if opt.AmazonEC2RootSize != 0 {
		args = append(args, "--amazonec2-root-size", fmt.Sprintf("%d", opt.AmazonEC2RootSize))
	}
	if opt.AmazonEC2RequestSpotInstance {
		args = append(args, "--amazonec2-request-spot-instance")
	}

	// Modify machine name to valid one
	if !AWSMachineNameExpression.MatchString(opt.Name) {
		opt.Name = strings.Replace(opt.Name, "_", "-", 1)
	}

	// args = append(args,
	// // https://docs.docker.com/machine/drivers/aws/#vpc-connectivity
	// "--amazonec2-private-address-only",
	// "--amazonec2-use-private-address",
	// )

	return args
}
