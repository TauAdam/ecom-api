package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CorrectPassword(hashedPassword string, plaintext []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), plaintext)
	return err == nil
}
