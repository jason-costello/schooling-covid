package services

import (
	"github.com/jason-costello/schooling-covid/internal/repositories"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type DistrictService interface {
}

type districtService struct {
	schoolRepo   *repositories.SchoolRepository
	districtRepo *repositories.DistrictRepository
	countRepo    *repositories.CountRepository
	logger       *logrus.Logger
}

func NewDistrictService(schoolRepo *repositories.SchoolRepository, countRepo *repositories.CountRepository, districtRepo *repositories.DistrictRepository, logger *log.Logger) *districtService {

	return &districtService{
		schoolRepo:   schoolRepo,
		districtRepo: districtRepo,
		countRepo:    countRepo,
		logger:       logger,
	}
}
