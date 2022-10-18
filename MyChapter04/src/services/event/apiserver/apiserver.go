package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"jaehonam.com/ev/config"
	"jaehonam.com/ev/database"
)

func Serve(c *config.ApiserverConfig, databaseHandler database.DatabaseHandler) (chan error, chan error) {
	handler := NewEventHandler(databaseHandler)

	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods(http.MethodGet).Path("/{criteria}/{key}").HandlerFunc(handler.FindEventHandler)
	eventsRouter.Methods(http.MethodGet).Path("").HandlerFunc(handler.AllEventHandler)
	eventsRouter.Methods(http.MethodPost).Path("").HandlerFunc(handler.NewEventHandler)

	httpErrChan := make(chan error)
	httpTLSErrChan := make(chan error)

	go func() {
		httpErrChan <- http.ListenAndServe(c.Endpoint, r)
	}()
	go func() {
		httpTLSErrChan <- http.ListenAndServeTLS(c.TLSEndpoint, "./cert.pem", "./key.pem", r)
	}()

	return httpErrChan, httpTLSErrChan
}
