package datastore_factory

import (
	"github.com/Kajekk/e-commerce-sys/order-service/config"
	"github.com/Kajekk/e-commerce-sys/order-service/container"
)

//type DatastoreFactory interface {
//	Build()
//}

//func init() {
//	//TODO inject specific data config to datastore map
//
//
//}

var DatastoreFactoryBuilderMap = map[string]DatastoreFbInterface{
	//
	config.MONGODB: &mongoDatastoreFactory{},
}

type DatastoreInterface interface{}

type DatastoreFbInterface interface {
	Build(c container.Container, appConfig *config.AppConfig) (DatastoreInterface, error)
}

func GetDatastoreFb(key string) DatastoreFbInterface {
	return DatastoreFactoryBuilderMap[key]
}
