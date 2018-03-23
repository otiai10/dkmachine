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
		IPAddress               string
		MachineName             string
		SSHUser                 string
		SSHPort                 int
		SSHKeyPath              string
		StorePath               string
		SwarmMaster             bool
		SwarmHost               string
		SwarmDiscovery          string
		ID                      string
		AccessKey               string
		SecretKey               string
		SessionToken            string
		Region                  string
		AMI                     string
		SSHKeyID                int
		ExistingKey             bool
		KeyName                 string
		InstanceID              string
		InstanceType            string
		PrivateIPAddress        string
		SecurityGroupID         string
		SecurityGroupIds        []string
		SecurityGroupName       string
		SecurityGroupNames      []string
		OpenPorts               interface{}
		Tags                    string
		ReservationID           string
		DeviceName              string
		RootSize                int
		VolumeType              string
		IamInstanceProfile      string
		VpcID                   string
		SubnetID                string
		Zone                    string
		RequestSpotInstance     bool
		SpotPrice               string
		BlockDurationMinutes    int
		PrivateIPOnly           bool
		UsePrivateIP            bool
		UseEbsOptimizedInstance bool
		Monitoring              bool
		SSHPrivateKeyPath       string
		RetryCount              int
		Endpoint                string
		DisableSSL              bool
		UserDataFile            string
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
