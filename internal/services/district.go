package services

import (
	"github.com/jason-costello/schooling-covid/internal/models"
	"github.com/jason-costello/schooling-covid/internal/repositories"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type DistrictService interface {

	Districts() ([]models.District, error)
	District(id int) (models.District, error)


}

type districtService struct {
	districtRepo *repositories.DistrictRepository
	logger       *logrus.Logger
}

func NewDistrictService( districtRepo *repositories.DistrictRepository, logger *log.Logger) *districtService {

	return &districtService{
		districtRepo: districtRepo,
		logger:       logger,
	}
}
func(d *districtService) Districts() ([]models.District, error){



	return nil, nil
}
func (d *districtService) District(id int) (models.District, error){

return models.District{}, nil
}
