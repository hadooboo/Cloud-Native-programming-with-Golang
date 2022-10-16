package mongodblayer

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"jaehonam.com/ev/config"
	"jaehonam.com/ev/database"
	"jaehonam.com/ev/model"
)

const (
	DB     = "ev"
	USERS  = "users"
	EVENTS = "events"
)

type MongoDBLayer struct {
	client *mongo.Client
}

func NewMongoDBLayer(c *config.DatabaseConfig) (database.DatabaseHandler, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.Conn))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return &MongoDBLayer{
		client: client,
	}, nil
}

func (ml *MongoDBLayer) AddEvent(e *model.Event) ([]byte, error) {
	if e.ID.IsZero() {
		e.ID = primitive.NewObjectID()
	}

	if e.Location.ID.IsZero() {
		e.Location.ID = primitive.NewObjectID()
	}

	_, err := ml.client.Database(DB).Collection(EVENTS).InsertOne(context.Background(), e)
	if err != nil {
		return nil, err
	}

	return e.ID[:], nil
}

func (ml *MongoDBLayer) FindEvent(id string) (*model.Event, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res := ml.client.Database(DB).Collection(EVENTS).FindOne(context.Background(), bson.D{{"_id", oid}})
	if err := res.Err(); err != nil {
		return nil, err
	}

	e := model.Event{}
	if err := res.Decode(&e); err != nil {
		return nil, err
	}

	return &e, nil
}

func (ml *MongoDBLayer) FindEventByName(name string) (*model.Event, error) {
	res := ml.client.Database(DB).Collection(EVENTS).FindOne(context.Background(), bson.D{{"name", name}})
	if err := res.Err(); err != nil {
		return nil, err
	}

	e := model.Event{}
	if err := res.Decode(&e); err != nil {
		return nil, err
	}

	return &e, nil
}

func (ml *MongoDBLayer) FindAllAvailableEvents() ([]*model.Event, error) {
	cur, err := ml.client.Database(DB).Collection(EVENTS).Find(context.Background(), bson.D{})
	if err != nil {
		println(err.Error())
		return nil, err
	}

	events := make([]*model.Event, 0)
	if err := cur.All(context.Background(), &events); err != nil {
		return nil, err
	}

	return events, nil
}
