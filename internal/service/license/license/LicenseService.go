package license

type LicenseService struct{}

func NewLicenseService() *LicenseService {
	return &LicenseService{}
}

func (s *LicenseService) List() []*license {
	return alllicenses
}

func (s *LicenseService) Get(idx int) (**license, error) {
	return &alllicenses[idx], nil
}

func (s *LicenseService) Delete(idx int) (bool, error) {
	copy(alllicenses[idx:], alllicenses[idx+1:])
	alllicenses[len(alllicenses)-1] = nil
	alllicenses = alllicenses[:len(alllicenses)-1]
	return true, nil
}
