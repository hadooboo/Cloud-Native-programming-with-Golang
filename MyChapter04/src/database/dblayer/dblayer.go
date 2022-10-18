package dblayer

import (
	"fmt"

	"jaehonam.com/ev/config"
	"jaehonam.com/ev/database"
	"jaehonam.com/ev/database/mongodblayer"
)

const (
	mongodb  string = "mongodb"
	dynamodb string = "dynamodb"
)

func NewDatabaseLayer(c *config.DatabaseConfig) (database.DatabaseHandler, error) {
	switch c.Type {
	case mongodb:
		return mongodblayer.NewMongoDBLayer(c)
	default:
		return nil, fmt.Errorf("invalid database type: %v", c.Type)
	}
}
