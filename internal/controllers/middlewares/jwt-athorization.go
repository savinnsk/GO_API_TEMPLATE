package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/savinnsk/api-template-go/configs"

)

type ctxKey string

const ClaimsKey ctxKey = "jwtClaims"

func JwtAuth(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		env , _ := configs.LoadEnv()

		tokenString := r.Header.Get("Authorization")

		
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
		}

		token, err := jwt.Parse(tokenString ,func(token *jwt.Token)(interface{},error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil,fmt.Errorf("athorization error")		
			}

			return []byte(env.JwtSecret), nil
		})

	
		if err != nil || !token.Valid || token == nil {
			log.Printf("JwtAuth token invalid: %v", err)
			http.Error(w,"athorization error",http.StatusUnauthorized)
			return
		}

		if claims , ok := token.Claims.(jwt.MapClaims) ; ok {
			ctx := context.WithValue(r.Context(),ClaimsKey, claims)
			next.ServeHTTP(w,r.WithContext(ctx))
			return
		} 

		http.Error(w, "athorization error", http.StatusUnauthorized)

	})
}