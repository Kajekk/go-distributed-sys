package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

// AppConfig represents the application config
type AppConfig struct {
	MongoDBConfig   DataStoreConfig `json:"mongodbConfig,omitempty"`
	CacheGrpcConfig DataStoreConfig `json:"cachegrpcConfig,omitempty"`
	//SQLConfig       DataStoreConfig `yaml:"sqlConfig"`
	//CouchdbConfig   DataStoreConfig `yaml:"couchdbConfig"`
	//CacheGrpcConfig DataStoreConfig `yaml:"cacheGrpcConfig"`
	//UserGrpcConfig  DataStoreConfig `yaml:"userGrpcConfig"`
	//ZapConfig       LogConfig       `yaml:"zapConfig"`
	//LorusConfig     LogConfig       `yaml:"logrusConfig"`
	//Log             LogConfig       `yaml:"logConfig"`
	UseCase     UseCaseConfig `json:"usecase,omitempty"`
	DataStoreDB []string      `json:"datastoredb,omitempty"`
}

type UseCaseConfig struct {
	//Registration RegistrationConfig `yaml:"registration"`
	ListOrder ListOrderConfig `json:"listOrder,omitempty"`
	//ListCourse   ListCourseConfig   `yaml:"listCourse"`
}

type ListOrderConfig struct {
	Code            string     `json:"code,omitempty"`
	OrderDataConfig DataConfig `json:"orderDataConfig,omitempty"`
	CacheDataConfig DataConfig `json:"cacheDataConfig,omitempty"`
}

type DataConfig struct {
	Code            string          `json:"code,omitempty"`
	DataStoreConfig DataStoreConfig `json:"dataStoreConfig,omitempty"`
}

type DataStoreConfig struct {
	Code string `json:"code,omitempty"`
	// Only database has a driver name, for grpc it is "tcp" ( network) for server
	DriverName string `json:"driverName,omitempty"`
	// For database, this is datasource name; for grpc, it is target url
	UrlAddress string `json:"urlAddress,omitempty"`
	// Only some databases need this database name
	DbName string `json:"dbName,omitempty"`
}

type LogConfig struct {
	// log library name
	Code string `yaml:"code"`
	// log level
	Level string `yaml:"level"`
	// show caller in log message
	EnableCaller bool `yaml:"enableCaller"`
}

const (
	DEVELOPMENT        = "development"
	PRODUCTION         = "production"
	CONFIG_DEVELOPMENT = "config_development"
	CONFIG_PRODUCTION  = "config_production"

	ORDER      = "order"
	ORDER_DATA = "orderData"

	MONGODB = "mongoDB"
)

func ReadConfig(env string) (*AppConfig, error) {
	var appConfig AppConfig
	var configStr string
	switch env {
	case DEVELOPMENT:
		configStr = os.Getenv(CONFIG_DEVELOPMENT)
	case PRODUCTION:
		configStr = os.Getenv(CONFIG_PRODUCTION)
	}
	decoded, err := base64.URLEncoding.DecodeString(configStr)
	if err != nil {
		fmt.Println("[Parse config] Convert B64 config string error: " + err.Error())
		return nil, errors.Wrap(err, "decode b64 config")
	}
	err = json.Unmarshal(decoded, &appConfig)
	if err != nil {
		fmt.Println("[Parse config] Parse JSON with config string error: " + err.Error())
		return nil, errors.Wrap(err, "unmarshal config")
	}

	return &appConfig, nil
}

func SetupDatastoreConfig(ac *AppConfig) {
	for ds := range ac.DataStoreDB {

	}
}
