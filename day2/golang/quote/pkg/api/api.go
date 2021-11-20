package api

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"quote/pkg/services/status"
)

type Server struct {
	httpPort      int
	server        *http.Server
	wg            sync.WaitGroup
	statusService status.StatusService
}

func NewServer(httpPort int, statusService status.StatusService) *Server {

	router := mux.NewRouter()

	server := &http.Server{
		Addr:           ":" + strconv.Itoa(httpPort),
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1000000,
	}
	s := &Server{
		httpPort:      httpPort,
		server:        server,
		statusService: statusService,
	}

	router.HandleFunc("/", welcome).Methods(http.MethodGet)

	router.HandleFunc("/api/status", s.status).Methods(http.MethodGet)

	return s
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to REST API"))
	w.WriteHeader(http.StatusOK)
}

func (s *Server) Run() error {
	s.wg.Add(1)
	var err error
	go func() {
		err = s.server.ListenAndServe()
		if err == http.ErrServerClosed {
			err = nil
		}
		s.wg.Done()
	}()

	s.wg.Wait()
	return nil
}

func (s *Server) Close() {
	err := s.server.Close()
	if err != nil {
		logrus.Infof("Error closing server, error: %s", err.Error())
	}
}
