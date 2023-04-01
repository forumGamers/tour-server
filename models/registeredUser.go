package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisteredUser struct {
	Id 				primitive.ObjectID		`json:"id" bson:"id,omitempty"`
	TeamId			primitive.ObjectID		`json:"teamId" bson:"teamId,omitempty"`
	TourId			primitive.ObjectID		`json:"tour" bson:"tour,omitempty"`
	CreatedAt		time.Time				`json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt		time.Time				`json:"updatedAt" bson:"updatedAt,omitempty"`
}