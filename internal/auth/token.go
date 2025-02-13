package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
		"roles":   roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return signedToken, nil
}

func ValidateToken(token string, roles []string) error {
	fmt.Println("Token", token)
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}

	if !t.Valid {
		return errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token mapping with claims")
	}

	userRoles, ok := claims["roles"].([]interface{})
	if !ok {
		return errors.New("roles claim is not an array")
	}

	for _, role := range userRoles {
		for _, r := range roles {
			if role == r {
				return nil
			}
		}
	}

	return errors.New("User Doesnt have the required roles")

}
