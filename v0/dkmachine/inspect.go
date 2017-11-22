package dkmachine

import (
	"encoding/json"
	"os/exec"
)

// Inspection ...
type Inspection struct {
	ConfigVersion int    `json:"config_version"`
	Name          string `json:"Name"`
	DriverName    string `json:"driver_name"`

	Driver struct {
		IPAddress           string   `json:"ip_address"`
		MachineName         string   `json:"machine_name"`
		SSHUser             string   `json:"ssh_user"`
		SSHPort             int      `json:"ssh_port"`
		SSHKeyPath          string   `json:"ssh_key_path"`
		StorePath           string   `json:"store_path"`
		SwarmMaster         bool     `json:"swarm_master"`
		SwarmHost           string   `json:"swarm_host"`
		SwarmDiscovery      string   `json:"swarm_discovery"`
		VBoxManager         struct{} `json:"vbox_manager"`
		HostInterfaces      struct{} `json:"host_interface"`
		CPU                 int      `json:"cpu"`
		Memory              int      `json:"memory"`
		DiskSize            int      `json:"disk_size"`
		NatNicType          string   `json:"nat_nictype"`
		Boot2DockerURL      string   `json:"boot2docker_url"`
		Boot2DockerImportVM string   `json:"boot2docker_import_vm"`
		HostDNSResolver     bool     `json:"host_dns_resolver"`
		HostOnlyCIDR        string   `json:"hostonly_cidr"`
		HostOnlyNicType     string   `json:"hostonly_nictype"`
		HostOnlyPromiscMode string   `json:"hostonly_promisc_mode"`
		UIType              string   `json:"ui_type"`
		HostOnlyNoDHCP      bool     `json:"hostonly_no_dhcp"`
		NoShare             bool     `json:"no_share"`
		DNSProxy            bool     `json:"dns_proxy"`
		NoVTXCheck          bool     `json:"no_vtx_check"`
		ShareFolder         string   `json:"share_folder"`
	} `json:"driver"`

	HostOptions struct {
		Driver string `json:"driver"`
		Memory int    `json:"memory"`
		Disk   int    `json:"disk"`

		EngineOptions struct {
			ArbitraryFlags   []interface{} `json:"arbitrary_flags"`
			DNS              interface{}   `json:"dns"`
			GraphDir         string        `json:"graph_dir"`
			Env              []interface{} `json:"env"`
			Ipv6             bool          `json:"ipv6"`
			InsecureRegistry []interface{} `json:"insecure_registry"`
			Labels           []interface{} `json:"labels"`
			LogLevel         string        `json:"log_level"`
			StorageDriver    string        `json:"storage_driver"`
			SelinuxEnabled   bool          `json:"selinux_enabled"`
			TLSVerify        bool          `json:"tls_verify"`
			RegistryMirror   []interface{} `json:"registry_mirror"`
			InstallURL       string        `json:"install_url"`
		} `json:"engine_options"`

		SwarmOptions struct {
			IsSwarm            bool          `json:"is_swarm"`
			Address            string        `json:"address"`
			Discovery          string        `json:"discovery"`
			Agent              bool          `json:"agent"`
			Master             bool          `json:"master"`
			Host               string        `json:"host"`
			Image              string        `json:"image"`
			Strategy           string        `json:"strategy"`
			Heartbeat          int           `json:"heartbeat"`
			Overcommit         int           `json:"overcommit"`
			ArbitraryFlags     []interface{} `json:"arbitrary_flags"`
			ArbitraryJoinFlags []interface{} `json:"arbitrary_join_flags"`
			Env                interface{}   `json:"env"`
			IsExperimental     bool          `json:"is_experimental"`
		} `json:"swarm_options"`

		AuthOptions struct {
			CertDir              string        `json:"cert_dir"`
			CaCertPath           string        `json:"ca_cert_path"`
			CaPrivateKeyPath     string        `json:"ca_private_key_path"`
			CaCertRemotePath     string        `json:"ca_cert_remote_path"`
			ServerCertPath       string        `json:"server_cert_path"`
			ServerKeyPath        string        `json:"server_key_path"`
			ClientKeyPath        string        `json:"client_key_path"`
			ServerCertRemotePath string        `json:"server_cert_remote_path"`
			ServerKeyRemotePath  string        `json:"server_key_remote_path"`
			ClientCertPath       string        `json:"client_cert_path"`
			ServerCertSANs       []interface{} `json:"server_cert_sans"`
			StorePath            string        `json:"store_path"`
		} `json:"auth_options"`
	} `json:"host_options"`
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
