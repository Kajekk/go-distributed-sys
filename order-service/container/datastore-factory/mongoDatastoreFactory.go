package datastore_factory

import (
	"context"
	"github.com/Kajekk/e-commerce-sys/order-service/config"
	"github.com/Kajekk/e-commerce-sys/order-service/container"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type mongoDatastoreFactory struct{}

func (m *mongoDatastoreFactory) Build(c container.Container, appConfig *config.AppConfig) (DatastoreInterface, error) {
	//TODO Setup DB Init here
	key := appConfig.MongoDBConfig.Code
	if value, ok := c.Get(key); ok {
		sdb := value.(*mongo.Client)
		return sdb, nil
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	var err error
	clientDB, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Check the connection
	err = clientDB.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	c.Put(appConfig.MongoDBConfig.Code, clientDB)
	//clientDB.Database(datastoreConfig.DbName)

	return clientDB, nil
}
