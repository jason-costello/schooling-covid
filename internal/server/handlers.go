package server

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func (s *Server) GetDistricts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Second)
	defer cancel()

	s.Logger.Debug("Getting districts")

	districts, err := s.SchoolDistrict.GetDistrictNames(ctx)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	data, err := json.Marshal(districts)
	if err != nil {
		s.Logger.Error(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(200)
	w.Write(data)
}

func (s *Server) GetDistrictSchools(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	ids := vars["did"]

	d, err := s.SchoolDistrict.AllSchoolsForDistrict(ctx, ids)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	//d.Schools = schools

	data, err := json.Marshal(d)
	if err != nil {
		s.Logger.Error(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}
