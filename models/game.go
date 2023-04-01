package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	Id 				primitive.ObjectID			`json:"id" bson:"id,omitempty"`
	Name			string						`json:"name" bson:"name,omitempty"`
	Type			string						`json:"type" bson:"type,omitempty"`  //solo / multiplayer
	Image			string						`json:"image" bson:"image,omitempty"`
	ImageId			string						`json:"imageId" bson:"imageId,omitempty"`
	Description		string						`json:"description" bson:"description"`
	CreatedAt		time.Time				`json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt		time.Time				`json:"updatedAt" bson:"updatedAt,omitempty"`
}