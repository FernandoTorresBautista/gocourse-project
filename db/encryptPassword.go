package db

import "golang.org/x/crypto/bcrypt"

// EncryptPassword ...
func EncryptPassword(password string) (string, error) {
	cost := 8 // 2 elevado al costo(2^costo) es el numero de pasadas sobre el password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
