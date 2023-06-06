package validation

import (
	"errors"
	"strconv"
	"time"

	h "github.com/forumGamers/tour-service/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateCreateTour(
	host, 
	name, 
	gameId, 
	pricePool, 
	slots, 
	startDate, 
	registrationFee, 
	location, 
	description, 
	tags string,
	) (primitive.ObjectID,int,int,time.Time,int,error) {
		_,err := primitive.ObjectIDFromHex(host)

		if err != nil {
			return primitive.ObjectID{},0,0,time.Time{},0,errors.New("Invalid ObjectID")
		}

		pool,err := strconv.ParseInt(pricePool,10,64)

		if err != nil {
			return primitive.ObjectID{},0,0,time.Time{},0,errors.New("Invalid data")
		}

		slot,err := strconv.ParseInt(slots,10,64)

		if err != nil {
			return primitive.ObjectID{},0,0,time.Time{},0,errors.New("Invalid data")
		}

		date,err := time.Parse("02-01-2006",startDate)

		if err != nil {
			return primitive.ObjectID{},0,0,time.Time{},0,errors.New("Invalid data")
		}

		fee,err := strconv.ParseInt(registrationFee,10,64)

		if err != nil {
			return primitive.ObjectID{},0,0,time.Time{},0,errors.New("Invalid data")
		}

		if err := h.ValidateInput(map[string]string{
			"name":name,
			"location":location,
			"description":description,
			"tags":tags,
		}) ; err != nil {
			return primitive.ObjectID{},0,0,time.Time{},0,errors.New("Invalid data")
		}

		id,err := primitive.ObjectIDFromHex(gameId)

		if err != nil {
			return primitive.ObjectID{},0,0,time.Time{},0,errors.New("Invalid data")
		}

		return id,int(pool),int(slot),date,int(fee),nil
}