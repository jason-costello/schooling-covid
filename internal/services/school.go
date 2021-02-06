package services

import (
	"github.com/jason-costello/schooling-covid/internal/repositories"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type SchoolService interface {
}

type schoolService struct {
	schoolRepo   *repositories.SchoolRepository
	districtRepo *repositories.DistrictRepository
	countRepo    *repositories.CountRepository
	logger       *logrus.Logger
}

func NewSchoolService(schoolRepo *repositories.SchoolRepository, countRepo *repositories.CountRepository, districtRepo *repositories.DistrictRepository, logger *log.Logger) *schoolService {

	return &schoolService{
		schoolRepo:   schoolRepo,
		districtRepo: districtRepo,
		countRepo:    countRepo,
		logger:       logger,
	}
}
