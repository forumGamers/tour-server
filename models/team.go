package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	Id 				primitive.ObjectID		`json:"id" bson:"id,omitempty"`
	Name			string					`json:"name" bson:"name,omitempty"`
	UserId			[]int					`json:"userId" bson:"userId,omitempty"`
	OwnerId			int						`json:"ownerId" bson:"ownerId,omitempty"`
	Image			string					`json:"image" bson:"image"`
	ImageId			string					`json:"imageId" bson:"imageId"`
	CreatedAt		time.Time				`json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt		time.Time				`json:"updatedAt" bson:"updatedAt,omitempty"`
	Description		string					`json:"description" bson:"description"`
}