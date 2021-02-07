package server

import (
	"net/http"
)

func (s *Server) GetDistricts(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	w.Write([]byte("hi"))
}

func (s *Server) GetDistrict(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hi"))

}
func (s *Server) GetDistrictSchools(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hi"))
}
func (s *Server) GetDistrictSchool(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hi"))
}
