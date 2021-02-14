// Package services provides an interface to be used to collect the needed data from various sources.
package services

import (
	"context"
	"fmt"
	"github.com/jason-costello/schooling-covid/internal/models"
	"github.com/jason-costello/schooling-covid/internal/repositories"
	"github.com/sirupsen/logrus"
)


type schoolDistrict struct {
	districtRepo repositories.District
	logger       *logrus.Logger
}

// SchoolDistrict interface allows access to data about school districts (districts, schools, school counts).
type SchoolDistrict interface {
	AllSchoolsForDistrict(ctx context.Context, districtShortName string) (map[string]models.School, error)
	GetDistrictNames(ctx context.Context) ([]models.District, error)
}


// NewSchoolDistrict returns a SchoolDistrict interface that has access to district info
func NewSchoolDistrict(districtsRepo repositories.District, logger *logrus.Logger) SchoolDistrict {

	return &schoolDistrict{
		districtRepo: districtsRepo,
		logger:       logger,
	}
}


// AllSchoolsForDistrict returns a map of school models.  The key for the map is the short-name of the school.  Counts
// for each school are also returned in this dataset.  The shortname of the district is captured in each school model
func (s *schoolDistrict) AllSchoolsForDistrict(ctx context.Context, districtShortName string) (map[string]models.School, error) {

	districtName, err := s.districtRepo.GetName(ctx, districtShortName)
	if err != nil {
		return map[string]models.School{}, err
	}
	fmt.Println("districtName: ", districtName)
	districtSchool, _ := s.districtRepo.DistrictByShortName(ctx, districtShortName)

	return districtSchool, nil

}


// GetDistrictNames returns a slice of districts.  For each district, the schools property will be nil in the returned
// dataset.  Only the name and short names will be populated.
func (s *schoolDistrict) GetDistrictNames(ctx context.Context) ([]models.District, error) {

	return s.districtRepo.AllDistricts(ctx)

}
