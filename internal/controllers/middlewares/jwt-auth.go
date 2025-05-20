package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/savinnsk/api-template-go/configs"

)

type ctxKey string

const ClaimsKey ctxKey = "jwtClaims"

func JwtAuth(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		tokenString := r.Header.Get("Authorization")

		if len(tokenString) > 7 && tokenString[:7] == "Bearer" {
			tokenString = tokenString[7:]
		}

		token, err := jwt.Parse(tokenString ,func(token *jwt.Token)(interface{},error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil,fmt.Errorf("unexpected signing method")		
			}

			env , _ := configs.LoadEnv()

			return env.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w,"Invalid token",http.StatusUnauthorized)
		}

		if claims , ok := token.Claims.(jwt.MapClaims) ; ok {
			ctx := context.WithValue(r.Context(),ClaimsKey, claims)
			next.ServeHTTP(w,r.WithContext(ctx))
			return
		} 

		http.Error(w, "Invalid claims", http.StatusUnauthorized)

	})
}