package dkmachine

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/otiai10/mint"
)

func TestMachine_RegenerateCerts(t *testing.T) {
	opt := &CreateOptions{
		Driver: "virtualbox",
		Name:   fmt.Sprintf("test-vb-%s", time.Now().Format("20060102150405")),
	}
	m, err := Create(opt)
	Expect(t, err).ToBe(nil)
	Expect(t, m).TypeOf("*dkmachine.Machine")
	defer m.Remove()

	ca := filepath.Join(m.CertPath(), "ca.pem")
	info, err := os.Stat(ca)
	Expect(t, err).ToBe(nil)

	original := info.ModTime().Unix()
	time.Sleep(1 * time.Second)

	err = m.RegenerateCerts()
	Expect(t, err).ToBe(nil)

	info, err = os.Stat(ca)
	Expect(t, err).ToBe(nil)
	regenerated := info.ModTime().Unix()

	Expect(t, regenerated).Not().ToBe(original)
}
