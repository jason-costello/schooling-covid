package services

import (
	"github.com/jason-costello/schooling-covid/internal/models"
	"github.com/jason-costello/schooling-covid/internal/repositories"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type CountService interface {
	Add(count models.Count, schoolID int) error
	Update(count models.Count) error
	Get(countID int) (models.Count, error)
}

type countService struct {
	schoolRepo   *repositories.SchoolRepository
	districtRepo *repositories.DistrictRepository
	countRepo    *repositories.CountRepository
	logger       *logrus.Logger
}

func NewCountService(countRepo *repositories.CountRepository,  logger *log.Logger) *countService {

	return &countService{
		countRepo:    countRepo,
		logger:       logger,
	}
}

func (c *countService) Add(count models.Count, schoolID int) error {
	return nil
}
func (c *countService) Update(count models.Count) error {
	return nil
}
func (c *countService) Get(countID int) (models.Count, error) {
	return models.Count{}, nil
}
