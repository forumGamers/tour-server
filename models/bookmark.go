package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bookmark struct {
	Id 			primitive.ObjectID		`json:"id" bson:"id,omitempty"`
	UserId		int						`json:"userId" bson:"userId"`
	TourId		primitive.ObjectID		`json:"tour" bson:"tour"`
	CreatedAt	time.Time				`json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt	time.Time				`json:"updatedAt" bson:"updatedAt,omitempty"`
}