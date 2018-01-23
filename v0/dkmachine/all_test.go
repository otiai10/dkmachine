package dkmachine

import (
	"strings"
	"testing"

	. "github.com/otiai10/mint"
)

func TestCreate(t *testing.T) {

	machine, err := Create(&CreateOptions{Name: "foo", Dry: true})

	Expect(t, err).ToBe(nil)
	Expect(t, machine).TypeOf("*dkmachine.Machine")
	Expect(t, strings.HasPrefix(machine.Name, "foo-")).ToBe(true)
	// jsonindent.NewEncoder(os.Stdout).Encode(machine)

	machine.Remove()
}

func TestCreate_aws(t *testing.T) {

	machine, err := Create(&CreateOptions{
		Name:   "foo_bar",
		Driver: "amazonec2",
		Dry:    true,
	})

	Expect(t, err).ToBe(nil)
	Expect(t, strings.HasPrefix(machine.Name, "foo-bar-")).ToBe(true)
	// jsonindent.NewEncoder(os.Stdout).Encode(machine)

	machine.Remove()
}
