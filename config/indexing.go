package config

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Indexing() {

	if _, err := Db.Collection("achievement").Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys: bson.M{
			"name":1,
		},
		Options: options.Index().SetUnique(true).SetBackground(true).SetName("achievement_name"),
	}) ; err != nil {
		panic(err.Error())
	}

	if _,err := Db.Collection("bookmark").Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys: bson.M{
			"userId":1,
		},
		Options: options.Index().SetUnique(true).SetBackground(true).SetName("bookmark_user"),
	}) ; err != nil {
		panic(err.Error())
	}

	if _,err := Db.Collection("champion").Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys: bson.M{
			"gameId":1,
		},
		Options: options.Index().SetUnique(true).SetBackground(true).SetName("champion_game_id"),
	}) ; err != nil {
		panic(err.Error())
	}

	if _,err := Db.Collection("champion").Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys: bson.M{
			"team":1,
		},
		Options: options.Index().SetUnique(true).SetBackground(true).SetName("champion_team"),
	}) ; err != nil {
		panic(err.Error())
	}

	if _,err := Db.Collection("game").Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys: bson.M{
			"name":1,
		},
		Options: options.Index().SetUnique(true).SetBackground(true).SetName("game_name"),
	}) ; err != nil {
		panic(err.Error())
	}

	if _,err := Db.Collection("registeredUser").Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys: bson.M{
			"tourId":1,
		},
		Options: options.Index().SetUnique(true).SetBackground(true).SetName("registered_user_tour"),
	}) ; err != nil {
		panic(err.Error())
	}

	if _,err := Db.Collection("team").Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys: bson.M{
			"name":1,
		},
		Options: options.Index().SetUnique(true).SetBackground(true).SetName("team_name"),
	}) ; err != nil {
		panic(err.Error())
	}

	if _,err := Db.Collection("teamAchievement").Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys: bson.M{
			"teamId":1,
		},
		Options: options.Index().SetUnique(true).SetBackground(true).SetName("team_achievement"),
	}) ; err != nil {
		panic(err.Error())
	}

	if _,err := Db.Collection("tour").Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys: bson.M{
			"host":1,
		},
		Options: options.Index().SetBackground(true).SetName("tour_host"),
	}) ; err != nil {
		panic(err.Error())
	}

	if _,err := Db.Collection("tour").Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys: bson.M{
			"gameId":1,
		},
		Options: options.Index().SetBackground(true).SetName("tour_game_id"),
	}) ; err != nil {
		panic(err.Error())
	}
}