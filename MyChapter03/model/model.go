package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	First    string
	Last     string
	Age      int
	Bookings []Booking
}

func (u *User) String() string {
	return fmt.Sprintf("id: %s, first_name: %s, last_name: %s, Age: %d, Bookings: %v", u.ID, u.First, u.Last, u.Age, u.Bookings)
}

type Booking struct {
	Date    int64
	EventID []byte
	Seats   int
}

type Event struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string
	Duration  int
	StartDate int64
	EndDate   int64
	Location  Location
}

type Location struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}
