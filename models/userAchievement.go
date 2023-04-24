package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAchievement struct {
	Id 					primitive.ObjectID		`json:"id" bson:"id,omitempty"`
	UserId				int						`json:"userId" bson:"userId,omitempty"`
	AchievementId		primitive.ObjectID		`json:"achievementId" bson:"achievementId,omitempty"`		
	CreatedAt			time.Time				`json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt			time.Time				`json:"updatedAt" bson:"updatedAt,omitempty"`
}