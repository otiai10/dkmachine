package dkmachine

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	// GoogleMachineNameExpression ...
	GoogleMachineNameExpression = regexp.MustCompile("^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$")
)

// ArgsForGoogleComputeEngine ...
func (opt *CreateOptions) ArgsForGoogleComputeEngine() []string {
	args := []string{
		"--driver", "google",
	}

	if opt.GoogleProject != "" {
		args = append(args, "--google-project", opt.GoogleProject)
	}

	if opt.GoogleZone != "" {
		args = append(args, "--google-zone", opt.GoogleZone)
	}

	if opt.GoogleScopes != "" {
		args = append(args, "--google-scopes", opt.GoogleScopes)
	}

	if opt.GoogleDiskSize != 0 {
		args = append(args, "--google-disk-size", fmt.Sprintf("%d", opt.GoogleDiskSize))
	}

	if !GoogleMachineNameExpression.MatchString(opt.Name) {
		// TODO: Refactor
		opt.Name = strings.Replace(opt.Name, ".", "-", -1)
		opt.Name = strings.Replace(opt.Name, "_", "-", -1)
		opt.Name = strings.ToLower(opt.Name)
	}

	return args
}
