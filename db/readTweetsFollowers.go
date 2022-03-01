package db

import (
	"context"
	"time"

	"github.com/FernandoTorresBautista/gocourse-project/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadTweetsFollowers ...
func ReadTweetsFollowers(ID string, page int) ([]models.TweetsFollowersBack, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("goproject")
	collection := db.Collection("relacion")

	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"usuarioid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"fecha": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	var result []models.TweetsFollowersBack
	cursor, err := collection.Aggregate(ctx, conditions)
	if err != nil {
		return nil, false
	}
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
