package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("Start app")

	return nil
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.configureRouter()

	s.logger.SetLevel(level)

	return http.ListenAndServe(s.config.AppPort, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handlerHallo())
}

func (s *APIServer) handlerHallo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "HELLO")
	}
}
