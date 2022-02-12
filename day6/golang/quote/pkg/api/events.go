package api

import (
	"encoding/json"
	"net/http"
	
	"github.com/go-playground/validator/v10"

	"github.com/sirupsen/logrus"

	"quote/pkg/models/dto"
)

func (s *Server) event(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	eventDto := dto.Event{}

	err := json.NewDecoder(r.Body).Decode(&eventDto)

	if err != nil {
		logrus.Errorf("error decoding json=%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	validate := validator.New();
	err = validate.Struct(eventDto)
	if err != nil {
		logrus.Errorf("Invalid post request=%v", err)
		w.WriteHeader((http.StatusBadRequest))
		return
	}

	logrus.WithFields(logrus.Fields{
		"day":         eventDto.Day,
		"month":       eventDto.Month,
		"year":        eventDto.Year,
		"title":       eventDto.Title,
		"description": eventDto.Description,
	}).Info("event Details")
	
	eventDto.Id = 1

	response, err := json.Marshal(eventDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(response))
	w.WriteHeader(http.StatusOK)
}
