package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"jaehonam.com/ev/config"
	"jaehonam.com/ev/database"
)

func Serve(c *config.ApiserverConfig, databaseHandler database.DatabaseHandler) error {
	handler := NewEventHandler(databaseHandler)

	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods(http.MethodGet).Path("/{criteria}/{key}").HandlerFunc(handler.FindEventHandler)
	eventsRouter.Methods(http.MethodGet).Path("").HandlerFunc(handler.AllEventHandler)
	eventsRouter.Methods(http.MethodPost).Path("").HandlerFunc(handler.NewEventHandler)
	return http.ListenAndServe(c.Endpoint, r)
}
