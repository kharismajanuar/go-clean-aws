package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PassBcrypt(pass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error generate password bcrypt", err)
		return "", err
	}
	return string(hashed), nil
}

func PassCompare(dbPass, inputPass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(inputPass)); err != nil {
		log.Println("error compare password bcrypt ~", err)
		return err
	}
	return nil
}
