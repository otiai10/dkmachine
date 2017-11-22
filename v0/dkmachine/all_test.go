package dkmachine

import (
	"os"
	"strings"
	"testing"

	"github.com/otiai10/jsonindent"
	. "github.com/otiai10/mint"
)

func TestCreate(t *testing.T) {

	machine, err := Create(&CreateOptions{Name: "foo", Dry: true})

	Expect(t, err).ToBe(nil)
	Expect(t, machine).TypeOf("*dkmachine.Machine")
	Expect(t, strings.HasPrefix(machine.Name, "foo-")).ToBe(true)

	jsonindent.NewEncoder(os.Stdout).Encode(machine)

	machine.Remove()
}
