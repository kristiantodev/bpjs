package confiq

import (
	"bpjs/dto/out"
	"bpjs/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtKey = []byte("bpjs")

type Claims struct {
	Id       int64   `json:"id"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	jwt.StandardClaims
}

func GenerateToken(user model.UserModel) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Id:       user.ID.Int64,
		Username: user.Username.String,
		Name:     user.FirstName.String+" "+user.LastName.String,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		_, err := ValidateToken(tokenString)
		if err != nil {
			out.ResponseOut(w, nil, false, http.StatusUnauthorized, "Unauthorized")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func DecodeToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}