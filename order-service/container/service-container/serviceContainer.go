package service_container

import (
	"github.com/Kajekk/e-commerce-sys/order-service/config"
	datastore_factory "github.com/Kajekk/e-commerce-sys/order-service/container/datastore-factory"
	"github.com/pkg/errors"
)

type ServiceContainer struct {
	FactoryMap map[string]interface{}
	AppConfig  *config.AppConfig
}

func (sc *ServiceContainer) InitApp(env string) error {
	var err error
	configApp, err := loadConfig(env)
	if err != nil {
		return errors.Wrap(err, "loadConfig")
	}
	sc.AppConfig = configApp

	return nil
}

func (sc *ServiceContainer) BuildUseCase(code string) (interface{}, error) {
	return nil, nil
}

func (sc *ServiceContainer) Get(code string) (interface{}, bool) {
	value, found := sc.FactoryMap[code]
	return value, found
}

func (sc *ServiceContainer) Put(code string, value interface{}) {
	sc.FactoryMap[code] = value
}

func (sc *ServiceContainer) BuildDataService(key string) (interface{}, error) {
	_, _ = datastore_factory.GetDatastoreFb(key).Build(sc, sc.AppConfig)
	return nil, nil
}

func loadConfig(env string) (*config.AppConfig, error) {
	ac, err := config.ReadConfig(env)
	if err != nil {
		return nil, errors.Wrap(err, "read container")
	}
	return ac, nil
}
