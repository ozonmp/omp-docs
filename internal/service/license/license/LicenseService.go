package license

type LicenseService struct{}

func NewLicenseService() *LicenseService {
	return &LicenseService{}
}

func (s *LicenseService) List() []license {
	return alllicenses
}

func (s *LicenseService) Get(idx int) (*license, error) {
	return &alllicenses[idx], nil
}
