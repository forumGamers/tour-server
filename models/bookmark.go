package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bookmark struct {
	Id 		primitive.ObjectID	`json:"id" bson:"id,omitempty"`
	UserId	int					`json:"userId" bson:"userId"`
	Tour	primitive.ObjectID	`json:"tour" bson:"tour"`
}