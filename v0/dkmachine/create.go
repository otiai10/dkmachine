package dkmachine

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os/exec"

	"github.com/google/uuid"
)

// CreateOptions ...
type CreateOptions struct {
	Name string `json:"name"`

	Help bool `json:"help"`
	Dry  bool `json:"dry"`

	// Name   string `json:"name"`
	Driver string `json:"driver"`

	EngineEnv              string `json:"engine_env"`
	EngineInsecureRegistry string `json:"engine_insecure_registry"`
	EngineInstallURL       string `json:"engine_install_url"`
	EngineLabel            string `json:"engine_label"`
	EngineOpt              string `json:"engine_opt"`
	EngineRegistryMirror   string `json:"engine_registry_mirror"`
	EngineStorageDriver    string `json:"engine_storage_driver"`

	Swarm             bool   `json:"swarm"`
	SwarmAddr         string `json:"swarm_addr"`
	SwarmDiscovery    bool   `json:"swarm_discovery"`
	SwarmExperimental bool   `json:"swarm_experimental"`
	SwarmHost         string `json:"swarm_host"`
	SwarmImage        string `json:"swarm_image"`
	SwarmMaster       bool   `json:"swarm_master"`
	SwarmOpt          string `json:"swarm_opt"`
	SwarmStrategy     string `json:"swarm_strategy"`

	TLSSan string `json:"tls_san"`

	// AmazonEC2 Options
	AmazonEC2AccessKey           string `json:"amazonec2_access_key"`
	AmazonEC2AMI                 string `json:"amazonec2_ami"`
	AmazonEC2InstanceType        string `json:"amazonec2_instance_type"`
	AmazonEC2Region              string `json:"amazonec2_region"`
	AmazonEC2IAMInstanceProfile  string `json:"amazonec2_iam_instance_profile"`
	AmazonEC2SecurityGroup       string `json:"amazonec2_security_group"`
	AmazonEC2RootSize            int    `json:"amazonec2_root_size"`
	AmazonEC2RequestSpotInstance bool   `json:"amazonec2_request_spot_instance"`

	// VirtualBox Options
	VirtualBoxBoot2DockerURL      string `json:"virtualbox_boot2docker_url"`
	VirtualBoxCPUCount            int    `json:"virtualbox_cpu_count"`
	VirtualBoxDiskSize            int    `json:"virtualbox_disk_size"`
	VirtualBoxHostDNSResolver     string `json:"virtualbox_host_dns_resolver"`
	VirtualBoxHostonlyCIDR        string `json:"virtualbox_hostonly_cidr"`
	VirtualBoxHostonlyNicPromisc  string `json:"virtualbox_hostonly_nicpromisc"`
	VirtualBoxHostonlyNicType     string `json:"virtualbox_hostonly_nictype"`
	VirtualBoxHostonlyNoDHCP      string `json:"virtualbox_hostonly_no_dhcp"`
	VirtualBoxImportBoot2DockerVM string `json:"virtualbox_import_boot2docker_vm"`
	VirtualBoxMemory              int    `json:"virtualbox_memory"`
	VirtualBoxNatNicType          string `json:"virtualbox_nat_nictype"`
	VirtualBoxNoDNSProxy          string `json:"virtualbox_no_dns_proxy"`
	VirtualBoxNoShare             string `json:"virtualbox_no_share"`
	VirtualBoxNoVTXCheck          string `json:"virtualbox_no_vtx_check"`
	VirtualBoxShareFolder         string `json:"virtualbox_share_folder"`
	VirtualBoxUIType              string `json:"virtualbox_ui_type"`
}

// Args ...
func (opt *CreateOptions) Args() []string {
	args := []string{"create"}
	if opt.Help {
		args = append(args, "--help")
		return args
	}

	switch opt.Driver {
	case "amazonec2":
		args = append(args, opt.ArgsForAmazonEC2()...)
	}
	opt.Name = id(opt.Name)
	args = append(args, opt.Name)
	return args
}

// Create ...
func Create(opt *CreateOptions) (*Machine, error) {

	if opt == nil {
		opt = &CreateOptions{}
	}
	args := opt.Args()

	machine := &Machine{
		CreateOptions: opt,
		Inspection:    &Inspection{Name: opt.Name},
	}

	if !opt.Dry {

		cmd := exec.Command(bin, args...)
		combinedoutput, err := cmd.CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("%v: %v", err, string(combinedoutput))
		}

		_, err = machine.Inspect()
		if err != nil {
			return nil, fmt.Errorf("failed to inspect created machine: %v", err)
		}

	}

	return machine, nil
}

func id(prefix ...string) string {
	prefix = append(prefix, "")
	if prefix[0] == "" {
		return genUUID()
	}
	return prefix[0] + "-" + genUUID()
}

func genUUID() string {
	buf := new(bytes.Buffer)
	uid := uuid.New()
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	defer encoder.Close()
	encoder.Write([]byte(uid.String()))
	return buf.String()[:8]
}
