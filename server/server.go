package server

import (
	"fmt"

	"github.com/premkit/premkit/handlers/v1"
	"github.com/premkit/premkit/log"
	"net/http"

	"github.com/gorilla/mux"
)

// Run is the main entrypoint of this daemon.
func Run(config *Config) {
	router := mux.NewRouter()

	internal := router.PathPrefix("/premkit").Subrouter()
	internalV1 := internal.PathPrefix("/v1").Subrouter()
	internalV1.HandleFunc("/service", v1.RegisterService).Methods("POST")

	// TODO serve the swagger.json using a gorilla static handlers

	forward := router.PathPrefix("/").Subrouter()
	forward.HandleFunc("/{path:.*}", v1.ForwardService)

	if config.HTTPPort != 0 {
		go func() {
			log.Infof("Listening on port %d for http connections", config.HTTPPort)
			log.Error(http.ListenAndServe(fmt.Sprintf(":%d", config.HTTPPort), router))
		}()
	}

	if config.HTTPSPort != 0 {
		go func() {
			log.Infof("Listening on port %d for https connections", config.HTTPSPort)
			log.Error(http.ListenAndServeTLS(fmt.Sprintf(":%d", config.HTTPSPort), config.TLSCertFile, config.TLSKeyFile, router))
		}()
	}
}
