package db

import (
	"context"
	"time"

	"github.com/FernandoTorresBautista/gocourse-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModifyRegister ....
func ModifyRegister(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("goproject")
	collection := db.Collection("usuarios")

	register := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		register["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		register["apellidos"] = u.Apellidos
	}
	register["fechaNacimiento"] = u.FechaNacimiento
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		register["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		register["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		register["sitioWeb"] = u.SitioWeb
	}

	updtString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := collection.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
