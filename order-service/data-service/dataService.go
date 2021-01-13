package data_service

import "github.com/Kajekk/e-commerce-sys/order-service/model"

type DataService interface {
	FindByField(interface{}) (*model.APIResponse, error)
}

type DataServiceImpl struct {
	DataServiceContainer interface{} //Data Connect
	TemplateObject       interface{} //Template Model
}
