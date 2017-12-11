package dkmachine

import (
	"encoding/json"
	"os/exec"
)

// Inspection ...
type Inspection struct {
	ConfigVersion int
	Name          string
	DriverName    string

	Driver struct {
		IPAddress           string
		MachineName         string
		SSHUser             string
		SSHPort             int
		SSHKeyPath          string
		StorePath           string
		SwarmMaster         bool
		SwarmHost           string
		SwarmDiscovery      string
		VBoxManager         struct{}
		HostInterfaces      struct{}
		CPU                 int
		Memory              int
		DiskSize            int
		NatNicType          string
		Boot2DockerURL      string
		Boot2DockerImportVM string
		HostDNSResolver     bool
		HostOnlyCIDR        string
		HostOnlyNicType     string
		HostOnlyPromiscMode string
		UIType              string
		HostOnlyNoDHCP      bool
		NoShare             bool
		DNSProxy            bool
		NoVTXCheck          bool
		ShareFolder         string
	}

	HostOptions struct {
		Driver string
		Memory int
		Disk   int

		EngineOptions struct {
			ArbitraryFlags   []interface{}
			DNS              interface{}
			GraphDir         string
			Env              []interface{}
			Ipv6             bool
			InsecureRegistry []interface{}
			Labels           []interface{}
			LogLevel         string
			StorageDriver    string
			SelinuxEnabled   bool
			TLSVerify        bool
			RegistryMirror   []interface{}
			InstallURL       string
		}

		SwarmOptions struct {
			IsSwarm            bool
			Address            string
			Discovery          string
			Agent              bool
			Master             bool
			Host               string
			Image              string
			Strategy           string
			Heartbeat          int
			Overcommit         int
			ArbitraryFlags     []interface{}
			ArbitraryJoinFlags []interface{}
			Env                interface{}
			IsExperimental     bool
		}

		AuthOptions struct {
			CertDir              string
			CaCertPath           string
			CaPrivateKeyPath     string
			CaCertRemotePath     string
			ServerCertPath       string
			ServerKeyPath        string
			ClientKeyPath        string
			ServerCertRemotePath string
			ServerKeyRemotePath  string
			ClientCertPath       string
			ServerCertSANs       []interface{}
			StorePath            string
		}
	}
}

// Inspect ...
func (m *Machine) Inspect() (*Inspection, error) {
	cmd := exec.Command(bin, "inspect", m.Name)
	o, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	inspection := new(Inspection)
	if err := json.Unmarshal(o, inspection); err != nil {
		return nil, err
	}
	m.Inspection = inspection
	return inspection, nil
}
