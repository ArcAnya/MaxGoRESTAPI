package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret" // should be harder to guess

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // expires in 2 hours
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(
		token,
		// Second argument: anonymous function called to validate the token
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC) // checking type of method: check if that token is signed with the correct method
			if !ok {
				return nil, errors.New("Invalid signing method")
			}
			return []byte(secretKey), nil
		})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	tokenIsValid := parsedToken.Valid // returns true if the token is valid

	if !tokenIsValid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims) // get hold of data stored in the token - checking that it is in the correct format of MapClaims

	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	// email := claims["email"].(string)  // getting the email from the token and casting it to a string
	userId := int64(claims["userId"].(float64)) // getting the userId from the token and casting it to an int64

	return userId, nil
}
