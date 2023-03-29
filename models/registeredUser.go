package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisteredUser struct {
	Id 				primitive.ObjectID		`json:"id" bson:"id,omitempty"`
	UserId			int						`json:"userId" bson:"userId,omitempty"`
	Tour			primitive.ObjectID		`json:"tour" bson:"tour,omitempty"`
	RegisteredAt	time.Time				`json:"registeredAt" bson:"registeredAt,omitempty"`
}