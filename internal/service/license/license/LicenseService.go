package license

import (
	"errors"
	"fmt"
	"github.com/istrel/bot/internal/model/license"
	"sort"
)

type LicenseService interface {
	Describe(licenseID uint32) (*license.License, error)
	List(cursor uint32, limit uint32) ([]license.License, error)
	Create(license.License) (uint32, error)
	Update(licenseID uint32, license license.License) error
	Remove(licenseID uint32) (bool, error)
}

type DummyLicenseService struct{}

func NewDummyLicenseService() *DummyLicenseService {
	return &DummyLicenseService{}
}

func (s *DummyLicenseService) Describe(licenseID uint32) (*license.License, error) {
	license, ok := alllicenses[licenseID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("License with ID %d is not found", licenseID))
	}

	return &license, nil
}

func (s *DummyLicenseService) List(cursor uint32, limit uint32) ([]license.License, error) {
	var retVal = make([]license.License, 0, len(alllicenses))

	for idx, value := range alllicenses {
		if idx > cursor && idx <= cursor+limit {
			retVal = append(retVal, value)
		}
	}

	sort.SliceStable(retVal, func(i, j int) bool {
		return retVal[i].ID < retVal[j].ID
	})

	return retVal, nil
}

func (s *DummyLicenseService) Create(license license.License) (uint32, error) {
	if _, ok := alllicenses[license.ID]; ok {
		return 0, errors.New(fmt.Sprintf("License with ID %d is exists", license.ID))
	}

	alllicenses[license.ID] = license
	return license.ID, nil
}

func (s *DummyLicenseService) Update(licenseID uint32, license license.License) error {
	if _, ok := alllicenses[licenseID]; !ok {
		return errors.New(fmt.Sprintf("License with ID %d is not exists", licenseID))
	}

	alllicenses[licenseID] = license
	return nil
}

func (s *DummyLicenseService) Remove(licenseID uint32) (bool, error) {
	if _, ok := alllicenses[licenseID]; !ok {
		return false, errors.New(fmt.Sprintf("License with ID %d is not found", licenseID))
	}

	delete(alllicenses, licenseID)
	return true, nil
}

func (s *DummyLicenseService) GetDataSize() uint32 {
	return uint32(len(alllicenses))
}
