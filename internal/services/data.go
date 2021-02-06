package services

import(
	"../storage/db"
	log "github.com/sirupsen/logrus"
)
func NewDataService(schoolRepo repositories.SchoolRepository, logger *log.Logger) *activationService {

	return &activationService{
		schoolRepository: schoolRepo,
		logger:         logger,
	}
}

type dataService struct {
	schoolRepo repositories.SchoolRepository
	logger         *log.Logger
}

