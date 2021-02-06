package server

import (
	"context"

	"github.com/jason-costello/schooling-covid/internal/repositories"
	"github.com/jason-costello/schooling-covid/internal/services"
	storage "github.com/jason-costello/schooling-covid/internal/storage/db"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"

)




type Application struct {
	CountSvc services.CountService
	SchoolSvc       services.SchoolService
	DistrictSvc   services.DistrictService
	WebServer     *http.Server
	Logger        *log.Logger
	DBConfig	DBConfig
	DB storage.DBTX
}


func NewApplication(dbConfig DBConfig, logger *logrus.Logger) Application{

	dbctx := context.Background()
	db, err := dbConfig.NewDB(dbctx)
	if err != nil{
		panic(err)
	}
	queries := storage.New(db)
	schoolRepo := repositories.NewSchoolRepository(queries, logger)
	countRepo := repositories.NewCountRepository(queries, logger)
	districtRepo := repositories.NewDistrictRepository(queries, logger)
	countSvc := services.NewCountService(&schoolRepo, &countRepo, &districtRepo, logger)
	schoolSvc := services.NewSchoolService(&schoolRepo, &countRepo, &districtRepo, logger)
	districtSvc := services.NewDistrictService(&schoolRepo, &countRepo, &districtRepo, logger)

	app :=  Application{
		CountSvc:    countSvc,
		SchoolSvc:   schoolSvc,
		DistrictSvc: districtSvc,
		WebServer:   nil,
		Logger:      nil,
		DB:   db,
	}
return app

}


func (app Application) SetRouter(r *mux.Router) {
	if r == nil {
		panic("no routes configured")
	}

	app.WebServer.Handler = r
}

func (app Application) Shutdown(c context.Context) {

	err := app.WebServer.Shutdown(c)
	if err != nil {

		app.Logger.Errorln(err)
	}

}

func (app Application) Serve() {
	// 	r := mux.NewRouter()

	go func() {

		if err := app.WebServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

}
func (app Application) ServeTLS() {
	// 	r := mux.NewRouter()

	go func() {
		if err := app.WebServer.ListenAndServeTLS("/webcerts/tls-cert.pem", "/webcerts/tls-key.pem"); err != nil {
			log.Println(err)
		}
	}()

}
