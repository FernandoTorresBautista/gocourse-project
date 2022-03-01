package db

import (
	"context"
	"time"

	"github.com/FernandoTorresBautista/gocourse-project/models"
)

// DeleteRelation ...
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("goproject")
	collection := db.Collection("relacion")

	_, err := collection.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
