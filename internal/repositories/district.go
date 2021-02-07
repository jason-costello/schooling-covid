package repositories

import (
	"context"
	"github.com/jason-costello/schooling-covid/internal/models"
	storage "github.com/jason-costello/schooling-covid/internal/storage/db"
	"github.com/sirupsen/logrus"
)

type DistrictRepository interface {
	Add(ctx context.Context, district models.District) error
	Update(ctx context.Context, district models.District) error
	Districts(ctx context.Context) ([]models.District, error)
	District(ctx context.Context, districtID int) (models.District, error)

}

func NewDistrictRepository(db *storage.Queries, logger *logrus.Logger) DistrictRepository {

	if db == nil {
		return nil
	}
	return &districtRepository{
		db:     db,
		logger: logger,
	}

}

type districtRepository struct {
	db     *storage.Queries
	logger *logrus.Logger
}

func (d *districtRepository) Add(ctx context.Context, district models.District) error {
	acd := storage.AddDistrictDataParams{
		Name:      district.Name,
		ShortName: district.ShortName,
	}
	return d.db.AddDistrictData(ctx, acd)
}
func (d *districtRepository) Update(ctx context.Context, district models.District) error {

	ucd := storage.UpdateDistrictDataParams{
		Name:      district.Name,
		ShortName: district.ShortName,
	}

	return d.db.UpdateDistrictData(ctx, ucd)

}

func (d *districtRepository) Districts(ctx context.Context) ([]models.District, error) {

	dist, err := d.db.GetAllDistricts(ctx)
	if err != nil {
		return nil, err
	}
	var allDistricts []models.District
	for _, dd := range dist {
		allDistricts = append(allDistricts, models.District{
			ID:        int(dd.ID),
			Name:      dd.Name,
			ShortName: dd.ShortName,
		})
	}
	return allDistricts, nil
}
func (d *districtRepository) District(ctx context.Context, districtID int) (models.District, error){

	dist, err := d.db.GetDistrictData(ctx, int32(districtID))
	if err != nil {
		return models.District{}, err
	}

	mc := models.District{
		ID:        int(dist.ID),
		Name:      dist.Name,
		ShortName: dist.ShortName,
	}
	return mc, nil

}
