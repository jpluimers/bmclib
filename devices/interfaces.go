package devices

import (
	"github.com/bmc-toolbox/bmclib/cfgresources"
)

// Bmc represents the requirement of items to be collected a server
type Bmc interface {
	ApplyCfg(*cfgresources.ResourcesConfig) error
	BiosVersion() (string, error)
	BmcType() string
	BmcVersion() (string, error)
	CPU() (string, int, int, int, error)
	Disks() ([]*Disk, error)
	IsBlade() (bool, error)
	License() (string, string, error)
	Login() error
	Logout() error
	Memory() (int, error)
	Model() (string, error)
	Name() (string, error)
	Nics() ([]*Nic, error)
	PowerKw() (float64, error)
	PowerState() (string, error)
	Serial() (string, error)
	Status() (string, error)
	TempC() (int, error)
	Vendor() string
	ServerSnapshot() (interface{}, error)
	UpdateCredentials(string, string)
}

// BmcChassis represents the requirement of items to be collected from a chassis
type BmcChassis interface {
	ApplyCfg(*cfgresources.ResourcesConfig) error
	Blades() ([]*Blade, error)
	BmcType() string
	ChassisSnapshot() (*Chassis, error)
	Close() error
	FindBladePosition(string) (int, error)
	FwVersion() (string, error)
	IsActive() bool
	IsOn() (bool, error)
	IsOnBlade(int) (bool, error)
	Model() (string, error)
	Name() (string, error)
	Nics() ([]*Nic, error)
	PassThru() (string, error)
	PowerCycle() (bool, error)
	PowerCycleBlade(int) (bool, error)
	PowerCycleBmcBlade(int) (bool, error)
	PowerKw() (float64, error)
	PowerOff() (bool, error)
	PowerOffBlade(int) (bool, error)
	PowerOn() (bool, error)
	PowerOnBlade(int) (bool, error)
	Psus() ([]*Psu, error)
	PxeOnceBlade(int) (bool, error)
	ReseatBlade(int) (bool, error)
	Serial() (string, error)
	SetFlexAddressState(int, bool) (bool, error)
	Status() (string, error)
	StorageBlades() ([]*StorageBlade, error)
	TempC() (int, error)
	UpdateCredentials(string, string)
	Vendor() string
}
