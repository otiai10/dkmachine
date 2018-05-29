package dkmachine

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/docker/machine/commands/mcndirs"
	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/auth"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/engine"
	"github.com/docker/machine/libmachine/host"
	"github.com/docker/machine/libmachine/swarm"

	"github.com/docker/machine/libmachine/drivers/rpc"
)

// Client ...
type Client struct {
	API           *libmachine.Client
	Host          *host.Host
	CreateOptions *CreateOptions
}

// NewClient ...
func NewClient(opt *CreateOptions) (*Client, error) {

	d := getDriver(opt)

	raw, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	api := libmachine.NewClient(
		mcndirs.GetBaseDir(),
		mcndirs.GetMachineCertDir(),
	)

	h, err := api.NewHost(opt.Driver, raw)
	if err != nil {
		return nil, err
	}

	name := d.GetMachineName()
	certdir := mcndirs.GetMachineCertDir()
	machinesdir := mcndirs.GetMachineDir()

	fmt.Println(
		"001",
		name,
		certdir,
		machinesdir,
	)

	h.HostOptions = &host.Options{
		AuthOptions: &auth.Options{
			CertDir:          certdir,
			CaCertPath:       filepath.Join(certdir, "ca.pem"),
			CaPrivateKeyPath: filepath.Join(certdir, "ca-key.pem"),
			ClientCertPath:   filepath.Join(certdir, "cert.pem"),
			ClientKeyPath:    filepath.Join(certdir, "key.pem"),
			ServerCertPath:   filepath.Join(machinesdir, name, "server.pem"),
			ServerKeyPath:    filepath.Join(machinesdir, name, "server-key.pem"),
			StorePath:        filepath.Join(machinesdir, name),
		},
		EngineOptions: &engine.Options{
			TLSVerify:  true,
			InstallURL: drivers.DefaultEngineInstallURL,
		},
		SwarmOptions: &swarm.Options{
			IsSwarm:   false,
			Master:    false,
			Discovery: "",
			Host:      "tcp://0.0.0.0:3376",
		},
	}

	return &Client{api, h, opt}, nil
}

// CreateFlags ...
func (client *Client) CreateFlags() rpcdriver.RPCFlags {
	machineflags := client.Host.Driver.GetCreateFlags()
	flags := rpcdriver.RPCFlags{
		Values: make(map[string]interface{}),
	}
	for _, f := range machineflags {
		flags.Values[f.String()] = f.Default()
		if f.Default() == nil {
			flags.Values[f.String()] = false
		}
	}
	return client.additionalOptionsForCreate(flags)
}

// AdditionalOptionsForCreate ...
func (client *Client) additionalOptionsForCreate(flags rpcdriver.RPCFlags) rpcdriver.RPCFlags {
	switch client.Host.DriverName {
	case "amazonec2":
		flags.Values["amazonec2-region"] = client.CreateOptions.AmazonEC2Region
	case "google":
	}
	return flags
}
