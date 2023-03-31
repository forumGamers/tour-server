package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tour struct {
	Id 					primitive.ObjectID		`json:"id" bson:"id,omitempty"`
	Name				string					`json:"name" bson:"name,omitempty"`
	GameId				primitive.ObjectID		`json:"gameId" bson:"gameId,omitempty"`
	PricePool			int						`json:"pricePool" bson:"pricePool,omitempty"`
	Slot				int						`json:"slot" bson:"slot,omitempty"`
	StartDate			time.Time				`json:"startDate" bson:"startDate,omitempty"`
	RegistrationFee		int						`json:"registrationFee" bson:"registrationFee,omitempty"`
	Status				string					`json:"status" bson:"status,omitempty"`
	Host				string					`json:"host" bson:"host,omitempty"`
	Location 			string					`json:"location" bson:"location,omitempty"`
	Description			string					`json:"description" bson:"description,omitempty"`
	Achievement			primitive.ObjectID		`json:"achievement" bson:"achievement,omitempty"`
	Tags				[]string				`json:"tags" bson:"tags"`
	Champion			[]primitive.ObjectID	`json:"champion" bson:"champion"`
	Image				string					`json:"image" bson:"image"`
	ImageId				string					`json:"imageId" bson:"imageId"`
	Game				Game
}