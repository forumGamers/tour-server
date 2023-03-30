package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Achievement struct {
	Id 			primitive.ObjectID		`json:"id" bson:"id,omitempty"`
	GameId 		primitive.ObjectID		`json:"gameId" bson:"gameId,omitempty"`
	Name		string					`json:"name" bson:"name,omitempty"`
	Image		string					`json:"image" bson:"image"`
}