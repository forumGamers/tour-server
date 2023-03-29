package loaders

import (
	cfg "github.com/forumGamers/tour-service/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDb() *mongo.Database {
	return cfg.Db
}