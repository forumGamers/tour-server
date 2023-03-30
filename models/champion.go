package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Champion struct {
	Id 			primitive.ObjectID		`json:"id" bson:"id,omitempty"`
	PlayerId	[]int					`json:"playerId" bson:"playerId,omitempty"`
	GameId		primitive.ObjectID		`json:"gameId" bson:"gameId,omitempty"`
	Team		primitive.ObjectID		`json:"team" bson:"team,omitempty"` //ambil dari service lain
	CreatedAt	time.Time				`json:"createdAt" bson:"createdAt"`
}