package db

import (
	"github.com/FernandoTorresBautista/gocourse-project/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin do the check of login in db
func TryLogin(email, password string) (models.Usuario, bool) {
	// check if user exist
	user, ok, _ := CheckIfUserExist(email)
	if !ok {
		return user, false
	}
	// check if password is equal
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
