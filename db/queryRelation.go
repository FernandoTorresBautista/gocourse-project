package db

import (
	"context"
	"time"

	// "fmt"

	"github.com/FernandoTorresBautista/gocourse-project/models"
	"go.mongodb.org/mongo-driver/bson"
)

// QueryRelation ...
func QueryRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("goproject")
	collection := db.Collection("relacion")

	condition := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var result models.Relation
	// fmt.Println(result)

	err := collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		// fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
