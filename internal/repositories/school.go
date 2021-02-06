package repositories

import (
	"context"
	"github.com/jason-costello/schooling-covid/internal/models"
	"github.com/jason-costello/schooling-covid/internal/storage/db"
	log "log"
)



type SchoolRepository interface{

	Add( ctx context.Context,school models.School, schoolID int) error
	Update( ctx context.Context,school models.School) error
	Get(ctx context.Context,schoolID int) (models.School,error)
}
func NewSchoolRepository(db  *storage.Queries, logger *log.Logger) *schoolRepository {

	if db == nil {
		return nil
	}
	return &schoolRepository{
		db:     db,
		logger: logger,
	}

}

type schoolRepository struct {
	db     *storage.Queries
	logger *log.Logger
}

func (c *schoolRepository) Add(ctx context.Context, school models.School, schoolID int)error{
	acd:=storage.AddSchoolDataParams{
		Name:    school.Name,
		ShortName: school.ShortName,
		DistrictID: int32(school.DistrictID),
	}
	return c.db.AddSchoolData(ctx, acd)
}
func (c *schoolRepository) Update(ctx context.Context, school models.School, schoolID int )error{

	ucd := storage.UpdateSchoolDataParams{
		Name:    school.Name,
		ShortName: school.ShortName,
		DistrictID: int32(school.DistrictID),
	}

	return c.db.UpdateSchoolData(ctx, ucd)


}
func (c *schoolRepository) Get(ctx context.Context,schoolID int) (models.School, error){


	d, err :=  c.db.GetSchoolData(ctx, int32(schoolID))
	if err != nil{
		return models.School{}, err
	}


	mc := models.School{
		ID:        int(d.ID),
		Name:      d.Name,
		ShortName: d.ShortName,
		DistrictID: int(d.DistrictID),
	}
	return mc, nil

}
