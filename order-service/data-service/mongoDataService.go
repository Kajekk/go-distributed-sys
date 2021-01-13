package data_service

import (
	"context"
	"github.com/Kajekk/e-commerce-sys/order-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

type mongoDataService struct {
	DbName string
	DataServiceImpl
}

var MongoService = &mongoDataService{}

func (m *mongoDataService) FindByField(field interface{}) *model.APIResponse {
	coll, ok := m.DataServiceImpl.DataServiceContainer.(*mongo.Collection)
	if !ok {
		return &model.APIResponse{}
	}

	obj, err := m.convertToBson(field)
	if err != nil {
		return &model.APIResponse{}
	}

	ctx := context.TODO()
	cur, _ := coll.Find(ctx, obj)

	list := m.NewList(100)
	_ = cur.All(ctx, &list)

	return &model.APIResponse{
		Data: list,
	}
}

func (m *mongoDataService) convertToBson(ent interface{}) (bson.M, error) {
	if ent == nil {
		return bson.M{}, nil
	}

	sel, err := bson.Marshal(ent)
	if err != nil {
		return nil, err
	}

	obj := bson.M{}
	_ = bson.Unmarshal(sel, &obj)

	return obj, nil
}

func (m *mongoDataService) NewList(limit int) interface{} {
	t := reflect.TypeOf(m.TemplateObject)
	return reflect.MakeSlice(reflect.SliceOf(t), 0, limit).Interface()
}
