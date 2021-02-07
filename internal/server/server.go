package server

import (
	"context"
	"fmt"

	"github.com/jason-costello/schooling-covid/internal/repositories"
	"github.com/jason-costello/schooling-covid/internal/services"
	storage "github.com/jason-costello/schooling-covid/internal/storage/db"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"

)




type Server struct {
	CountSvc services.CountService
	SchoolSvc       services.SchoolService
	DistrictSvc   services.DistrictService
	WebServer     *http.Server
	Logger        *log.Logger
	DBConfig	DBConfig
	DB storage.DBTX
	Port int
}


func NewServer(port int, dbConfig DBConfig, logger *logrus.Logger) *Server{

	dbctx := context.Background()
	db, err := dbConfig.NewDB(dbctx)
	if err != nil{
		panic(err)
	}

	queries := storage.New(db)
	schoolRepo := repositories.NewSchoolRepository(queries, logger)
	countRepo := repositories.NewCountRepository(queries, logger)
	districtRepo := repositories.NewDistrictRepository(queries, logger)
	countSvc := services.NewCountService(&countRepo, logger)
	schoolSvc := services.NewSchoolService(&schoolRepo, logger)
	districtSvc := services.NewDistrictService(&districtRepo, logger)

	s :=  &Server{
		CountSvc:    countSvc,
		SchoolSvc:   schoolSvc,
		DistrictSvc: districtSvc,
		WebServer:   &http.Server{},
		Logger:      logger,
		DB:   db,
		Port: port,
	}

return s

}


func (s *Server) SetRouter(r *mux.Router) {
	if r == nil {
		panic("no routes configured")
	}

	s.WebServer.Handler = r
}

func (s *Server) Shutdown(c context.Context) {

	err := s.WebServer.Shutdown(c)
	if err != nil {

		s.Logger.Errorln(err)
	}

}

func (s *Server) Serve() {
	// 	r := mux.NewRouter()
	var addr = fmt.Sprintf(":%d", s.Port)
	s.WebServer.Addr = addr

	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			if err := s.WebServer.ListenAndServe(); err != nil {
				s.Logger.Println(err)
			}
		}
		}()
		s.Logger.Info("Serving on ", addr)
	}

func (s *Server) ServeTLS() {
	// 	r := mux.NewRouter()

	go func() {
		if err := s.WebServer.ListenAndServeTLS("/webcerts/tls-cert.pem", "/webcerts/tls-key.pem"); err != nil {
			log.Println(err)
		}
	}()

}

