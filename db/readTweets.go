package db

import (
	"context"
	"log"
	"time"

	"github.com/FernandoTorresBautista/gocourse-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadTweets ...
func ReadTweets(ID string, page int64) ([]*models.TweetBack, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("goproject")
	collection := db.Collection("tweet")

	var result []*models.TweetBack

	condition := bson.M{
		"userid": ID,
	}
	optionsTweet := options.Find()
	optionsTweet.SetLimit(20)
	optionsTweet.SetSort(bson.D{{Key: "fecha", Value: -1}})
	optionsTweet.SetSkip((page - 1) * 20)

	cursor, err := collection.Find(ctx, condition, optionsTweet)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}
	for cursor.Next(context.TODO()) {
		var register models.TweetBack
		err := cursor.Decode(&register)
		if err != nil {
			return result, false
		}
		result = append(result, &register)
	}
	return result, true
}
