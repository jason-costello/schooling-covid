// Package repositories provides interfaces into the underlying data layer
package repositories

import (
	"context"
	"github.com/jason-costello/schooling-covid/internal/models"
	storage "github.com/jason-costello/schooling-covid/internal/storage/db"
	"github.com/sirupsen/logrus"
	"time"
)

// District allows access to the school district specific data.  Get a list of all school districts
// or return all schools and count data for one school district.
type District interface {
	AllDistricts(ctx context.Context) ([]models.District, error)
	DistrictByShortName(ctx context.Context, districtShortName string) (map[string]models.School, error)
	GetName(ctx context.Context, shortname string) (string, error)
}

// NewDistrictRepository returns an instance of the District Interface that provides
// access to the function needed to collect district level data
func NewDistrictRepository(db storage.DBTX, logger *logrus.Logger) District {


	queries := storage.New(db)

	if db == nil {
		return nil
	}
	return &districtRepository{
		db:     queries,
		logger: logger,
	}

}

type districtRepository struct {
	db     *storage.Queries
	logger *logrus.Logger
}


// AllDistricts returns a slice of every district including shortName and name.  Schools for each
// district in this return will all be nil.  To collect school information, use a district short name
// and query 'DistrictByShortName'
func (d *districtRepository) AllDistricts(ctx context.Context) ([]models.District, error) {

	var dm []models.District
	dist, err := d.db.GetAllDistricts(ctx)
	if err != nil {
		return nil, err
	}
	for _, dd := range dist {
		dm = append(dm, models.District{
			Name:      dd.Name,
			ShortName: dd.ShortName,
		})
	}
	return dm, nil
}

// DistrictByShortName will return all of the schools and their associated count data for the provided districtShortName.
func (d *districtRepository) DistrictByShortName(ctx context.Context, districtShortName string) (map[string]models.School, error) {

	schools, err := d.db.GetAllSchoolsForDistrict(ctx, districtShortName)
	if err != nil {
		return map[string]models.School{}, err
	}


	schoolMap := make(map[string]models.School)
	var current, last string
	var schoolCounts models.SchoolCounts
	for i, school := range schools {

		current = school.ShortName_2

		if current != last && i != 0 {
			//for x, c := range schoolCounts {
			//
			//	schoolMap[last].Counts[x] =schoolCounts[x]
			//		}
			d.logger.Infoln("Reset schoolCounts to empty")
			schoolCounts = *new(models.SchoolCounts)
		}

		tf, _ := time.Parse("2006-01-02", school.CountDate)

		schoolCounts = append(schoolCounts, models.SchoolCount{
			CountDate:   tf,
			Positive:    int(school.Positive),
			Symptomatic: int(school.Symptomatic),
			Exposed:     int(school.Exposed),
			SchoolSn:    school.ShortName_2,
		},
		)

		schoolMap[school.Name_2] = models.School{
			DistrictShortName: school.DistrictShortName,
			Name:              school.Name_2,
			ShortName:         school.ShortName_2,
			Counts:            schoolCounts,
		}


		last = current

	}

	return schoolMap, nil

}


// GetName will return the full district name when passed the shortname.
func (d *districtRepository) GetName(ctx context.Context, districtShortName string) (string, error) {

	if ctx == nil{
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
		defer cancel()
	dist, err := d.db.GetDistrict(ctx, districtShortName)
	if err != nil {
		return "", err
	}

	return dist.ShortName, nil

}
