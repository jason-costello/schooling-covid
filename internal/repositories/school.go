package repositories
import (
	"../storage/db"
	log "github.com/sirupsen/logrus"
)



type SchoolRepository interface{


}
func NewSchoolRepository(db *storage.Queries, logger *log.Logger) *schoolRepository {

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

