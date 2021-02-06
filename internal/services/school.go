//go:generate moq -out school_service_moq_test.go . SchoolService

package services

import (
	"../models"
	"../repositories"
	"../storage/db"
	"github.rackspace.com/SegmentSupport/antivirus/pkg/storage/db"
)

type SchoolService interface {


	GetAllCounts() []models.School

}
func NewActivationService(dbtx storage.DBTX, coreClient Core, skuRPMRepo repositories.SkuToRpmRepository, logger *log.Logger) *activationService {

	return &activationService{
		db: *storage.Queries,
		logger:         logger,
	}
}

type activationService struct {
	activationRepo repositories.ActivationRepository
	skuRPMRepo     repositories.SkuToRpmRepository
	coreClient     Core
	logger         *log.Logger
}
