package database

import "jaehonam.com/ev/model"

type DatabaseHandler interface {
	AddEvent(*model.Event) ([]byte, error)
	FindEvent(string) (*model.Event, error)
	FindEventByName(string) (*model.Event, error)
	FindAllAvailableEvents() ([]*model.Event, error)
}
