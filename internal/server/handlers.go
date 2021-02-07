package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetDistricts(w http.ResponseWriter, r *http.Request) {
	districts, err := s.DistrictSvc.Districts()
	if err != nil{
		w.WriteHeader(500)
		return
	}

	data, err := json.Marshal(districts)
	if err != nil{
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(data)
	return
}

func (s *Server) GetDistrict(w http.ResponseWriter, r *http.Request) {

}
func (s *Server) GetDistrictSchools(w http.ResponseWriter, r *http.Request) {

}
func (s *Server) GetDistrictSchool(w http.ResponseWriter, r *http.Request) {

}
