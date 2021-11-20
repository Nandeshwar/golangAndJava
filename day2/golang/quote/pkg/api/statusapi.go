package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) status(w http.ResponseWriter, r *http.Request) {
	appStatus := s.statusService.GetAppStatus()

	response, err := json.Marshal(appStatus)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(response))
	w.WriteHeader(http.StatusOK)
}
