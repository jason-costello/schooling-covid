package repositories

import (
	"context"
	"github.com/jason-costello/schooling-covid/internal/models"
	"github.com/jason-costello/schooling-covid/internal/storage/db"
	log "log"
)



type DistrictRepository interface{

	Add( ctx context.Context,district models.District, schoolID int) error
	Update( ctx context.Context,district models.District) error
	Get(ctx context.Context,districtID int) (models.District,error)
}
func NewDistrictRepository(db  *storage.Queries, logger *log.Logger) *districtRepository {

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
	logger *log.Logger
}

func (c *districtRepository) Add(ctx context.Context, district models.District, schoolID int)error{
	acd:=storage.AddDistrictDataParams{
		Name:    district.Name,
		ShortName: district.ShortName,
	}
	return c.db.AddDistrictData(ctx, acd)
}
func (c *districtRepository) Update(ctx context.Context, district models.District, schoolID int )error{

	ucd := storage.UpdateDistrictDataParams{
		Name:    district.Name,
		ShortName: district.ShortName,
	}

	return c.db.UpdateDistrictData(ctx, ucd)


}
func (c *districtRepository) Get(ctx context.Context,districtID int) (models.District, error){


	d, err :=  c.db.GetDistrictData(ctx, int32(districtID))
	if err != nil{
		return models.District{}, err
	}


	mc := models.District{
		ID:        int(d.ID),
		Name:      d.Name,
		ShortName: d.ShortName,
	}
	return mc, nil

}
