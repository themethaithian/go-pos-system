package authen

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	"github.com/themethaithian/go-pos-system/config"
)

type AccessTokenClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func CreateJWT(userID string, role string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := AccessTokenClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(config.Val.JWTSecret)
	if err != nil {
		return "", errors.Wrapf(err, "failed to sign token")
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*AccessTokenClaims, error) {
	claims := &AccessTokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return config.Val.JWTSecret, nil
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse token")
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
