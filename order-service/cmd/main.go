package main

import (
	"fmt"
	"github.com/Kajekk/e-commerce-sys/order-service/config"
	"github.com/Kajekk/e-commerce-sys/order-service/container"
	service_container "github.com/Kajekk/e-commerce-sys/order-service/container/service-container"
	"github.com/pkg/errors"
	"os"
)

func main() {
	//cf := os.Getenv("config")
	//
	//fmt.Println(cf)
	//
	//decoded, _ := base64.URLEncoding.DecodeString(cf)
	//var config1 config.AppConfig
	//
	//_ = json.Unmarshal(decoded, &config1)
	//fmt.Println(config1.MongoDBConfig.Code)
	env := os.Getenv("env")
	if env == "" {
		env = "stg"
	}
	appContainer, err := buildContainer(env)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Println(appContainer)

	//_ = setupHTTPServer(container)

}

func buildContainer(env string) (container.Container, error) {
	factoryMap := make(map[string]interface{})
	appConfig := config.AppConfig{}
	serviceContainer := service_container.ServiceContainer{
		FactoryMap: factoryMap,
		AppConfig:  &appConfig,
	}

	err := serviceContainer.InitApp(env)
	if err != nil {
		return nil, errors.Wrap(err, "InitApp")
	}

	return &serviceContainer, nil
}

//func setUpGrpcServer(sc *service_container.ServiceContainer) error {
//srv := grpc.NewServer()
//pb.RegisterRouteGuideServer(grpcServer, newServer())
//grpcServer.Serve(lis)

//	return nil
//}

//func setupHTTPServer(container container.Container) error {
//	return nil
//}

//func setupFactory(c container.Container) {
//	getOrderUseCase(c)
//}
//
//func getOrderUseCase(c container.Container) {
//	key := config.ORDER
//	_, _ = c.BuildUseCase(key)
//}
