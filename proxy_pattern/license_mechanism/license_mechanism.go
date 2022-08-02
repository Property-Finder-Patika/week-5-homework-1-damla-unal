package license_mechanism

import (
	"errors"
	client "proxy_pattern/client"
	sPack "proxy_pattern/software_package"
)

const licenseLimit = 10

type LicenseMechanism interface {
	GetLicense()
}

type LicenseProxy struct {
	Client client.Client
}

func (l LicenseProxy) GetLicense() (*sPack.SoftwarePackage, error) {
	sPackage := l.Client.LicensedPackage
	if sPackage.LicenseNumber >= licenseLimit {
		return nil, errors.New("all of the licenses have all been taken up and retry later")
	} else {
		sPackage.GiveNewLicense()
		return l.Client.LicensedPackage, nil
	}

}
