package jwt

import (
	"time"

	"github.com/FernandoTorresBautista/gocourse-project/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(t models.Usuario) (string, error) {
	myKey := []byte("key_of_jwt_to_sign_the_payload")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
