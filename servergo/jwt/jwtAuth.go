package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")
var emisor = "ingesis.uniquindio.edu.co"

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iss": emisor,
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func VerifyToken(tokenString string, username string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	// Verifica que la firma del token sea valido y la fecha de expiración
	if !token.Valid {
		return fmt.Errorf("token invalido | Firma o Fecha de expiración (exp) no valido")
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["iss"] != emisor {
		return fmt.Errorf("ISS (emisor) incorrecto")
	}
	fmt.Printf("sub: %s | username: %s\n", claims["sub"], username)
	if claims["sub"] != username {
		return fmt.Errorf("SUB (usuario) incorrecto")
	}

	return nil
}
