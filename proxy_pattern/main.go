package main

import (
	"fmt"
	"github.com/google/uuid"
	client "proxy_pattern/client"
	licence "proxy_pattern/license_mechanism"
	sPack "proxy_pattern/software_package"
)

func main() {
	desiredPackage := sPack.SoftwarePackage{LicenseNumber: 9, LicenseKey: uuid.New().String()}
	c := client.Client{LicensedPackage: &desiredPackage}

	licenseMechanism := licence.LicenseProxy{
		Client: c,
	}

	license, err := licenseMechanism.GetLicense()
	if err != nil {
		return
	}
	fmt.Printf("License generated for you. License Key: %s\n", license.LicenseKey)

	fmt.Println("------> Request a second license: ")
	license2, err := licenseMechanism.GetLicense()
	if err != nil {
		fmt.Printf("WARNING: %s\n", err)
		return
	}
	fmt.Printf("License generated for you. License Key: %s\n", license2.LicenseKey)
}
