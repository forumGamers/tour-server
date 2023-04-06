package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TeamAchievement struct {
	Id 				primitive.ObjectID		`json:"id" bson:"id,omitempty"`
	TeamId			primitive.ObjectID		`json:"teamId" bson:"teamId,omitempty"`
	AchievementId	primitive.ObjectID		`json:"achievementId" bson:"achievementId,omitempty"`		
}