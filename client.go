package dkmachine

import (
	"encoding/json"
	"path/filepath"
	"reflect"

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
		return client.additionalOptionsForCreateAmazonEC2(flags)
	case "google":
	}
	return flags
}

func (client *Client) additionalOptionsForCreateAmazonEC2(flags rpcdriver.RPCFlags) rpcdriver.RPCFlags {
	opt := client.CreateOptions
	values := flags.Values
	appendIfNotZero(values, opt.AmazonEC2Region, "amazonec2-region")
	appendIfNotZero(values, opt.AmazonEC2InstanceType, "amazonec2-instance-type")
	appendIfNotZero(values, opt.AmazonEC2IAMInstanceProfile, "amazonec2-iam-instance-profile")
	appendIfNotZero(values, []string{opt.AmazonEC2SecurityGroup}, "amazonec2-security-group")
	appendIfNotZero(values, opt.AmazonEC2RootSize, "amazonec2-root-size")
	appendIfNotZero(values, opt.AmazonEC2RequestSpotInstance, "amazonec2-request-spot-instance")

	// FIXME: opt.SwarmXxx
	values["swarm-master"] = false
	values["swarm-host"] = ""
	values["swarm-discovery"] = ""

	flags.Values = values
	return flags
}

func appendIfNotZero(values map[string]interface{}, value interface{}, name string) map[string]interface{} {
	if isZero(reflect.ValueOf(value)) {
		return values
	}
	values[name] = value
	return values
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}
