package repositories

import (
	"context"
	"github.com/jason-costello/schooling-covid/internal/models"
	"github.com/jason-costello/schooling-covid/internal/storage/db"
	log "log"
)



type CountRepository interface{

	Add( ctx context.Context,count models.Count, schoolID int) error
	Update( ctx context.Context,count models.Count) error
	Get(ctx context.Context,countID int) (models.Count,error)
}
func NewCountRepository(db  *storage.Queries, logger *log.Logger) *countRepository {

	if db == nil {
		return nil
	}
	return &countRepository{
		db:     db,
		logger: logger,
	}

}

type countRepository struct {
	db     *storage.Queries
	logger *log.Logger
}

func (c *countRepository) Add(ctx context.Context, count models.Count, schoolID int)error{
	acd:=storage.AddCountDataParams{
		SchoolID:    int32(schoolID),
		Positive:    int32(count.Positive),
		Symptomatic: int32(count.Symptomatic),
		Exposed:     int32(count.Exposed),
	}
	return c.db.AddCountData(ctx, acd)
}
func (c *countRepository) Update(ctx context.Context, count models.Count, schoolID int )error{

	ucd := storage.UpdateCountDataParams{
		SchoolID: int32(schoolID),
		Positive: int32(count.Positive),
		Exposed:  int32(count.Exposed),
		Symptomatic:       int32(count.Symptomatic),
		ID:  int32(count.ID),
	}

	return c.db.UpdateCountData(ctx, ucd)


}
func (c *countRepository) Get(ctx context.Context,countID int) (models.Count, error){


	cnt, err :=  c.db.GetCountData(ctx, int32(countID))
	if err != nil{
		return models.Count{}, err
	}

	mc := models.Count{
		ID:         int(cnt.ID),
		Observed:   cnt.CreatedAt,
		Positive:    int(cnt.Positive),
		Symptomatic: int(cnt.Symptomatic),
		Exposed:     int(cnt.Exposed),
	}
return mc, nil

}
