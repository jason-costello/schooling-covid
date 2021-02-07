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

func NewSchoolService(schoolRepo *repositories.SchoolRepository, logger *log.Logger) *schoolService {

	return &schoolService{
		schoolRepo:   schoolRepo,
		logger:       logger,
	}
}
