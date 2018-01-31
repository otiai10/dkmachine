package dkmachine

import (
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

	if !GoogleMachineNameExpression.MatchString(opt.Name) {
		// TODO: Refactor
		opt.Name = strings.Replace(opt.Name, ".", "-", -1)
		opt.Name = strings.Replace(opt.Name, "_", "-", -1)
		opt.Name = strings.ToLower(opt.Name)
	}

	return args
}
