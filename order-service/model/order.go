package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID          primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	CreatedTime time.Time          `json:"created_time,omitempty" bson:"created_time,omitempty"`
	OrderCode   string             `json:"order_code,omitempty" bson:"order_code,omitempty"`
}
