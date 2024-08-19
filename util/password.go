package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(pw string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func IsCorrectPassword(hashedPw, pw []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashedPw, []byte(pw)); err != nil {
		return false, err
	}

	return true, nil
}
