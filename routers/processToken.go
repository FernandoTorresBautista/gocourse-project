package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/FernandoTorresBautista/gocourse-project/db"
	"github.com/FernandoTorresBautista/gocourse-project/models"
)

// Email ...
var Email string

// IDUsuario ...
var IDUsuario string

// ProcessToken to extract the values
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("key_of_jwt_to_sign_the_payload")
	claims := &models.Claim{}
	splitToken := strings.Split("tk", "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, ok, ID := db.CheckIfUserExist(claims.Email)
		if ok {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, ok, ID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token inválido")
	}
	return claims, false, string(""), err
}
