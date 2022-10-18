package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"jaehonam.com/ev/database"
	"jaehonam.com/ev/model"
)

type eventHandler struct {
	databaseHandler database.DatabaseHandler
}

func NewEventHandler(dh database.DatabaseHandler) *eventHandler {
	return &eventHandler{
		databaseHandler: dh,
	}
}

func (eh *eventHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["criteria"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error": "No search criteria found, you can either
						search by id via /id/4
						to search by name via /name/coldplayconcert}"`)
		return
	}

	key, ok := vars["key"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error": "No search keys found, you can either search
						by id via /id/4
						to search by name via /name/coldplayconcert}"`)
		return
	}

	var event *model.Event
	var err error
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.databaseHandler.FindEventByName(key)
	case "id":
		event, err = eh.databaseHandler.FindEvent(key)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error": "criteria should be name or id"}`)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}

func (eh *eventHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	events, err := eh.databaseHandler.FindAllAvailableEvents()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Error occured while trying to find all available events: %s"}`, err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&events)
}

func (eh *eventHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {
	event := new(model.Event)
	err := json.NewDecoder(r.Body).Decode(&event)
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Error occured while decoding event data: %s"}`, err)
		return
	}

	id, err := eh.databaseHandler.AddEvent(event)
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Error occured while persisting event %d: %s"}`, id, err)
		return
	}

	fmt.Fprintf(w, `{"id": %d}`, id)
}
