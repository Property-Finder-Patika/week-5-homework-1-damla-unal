package software_package

type SoftwarePackage struct {
	LicenseNumber int
	LicenseKey    string
}

// GiveNewLicense increments the license number of the software package
func (s *SoftwarePackage) GiveNewLicense() {
	s.LicenseNumber++
}
