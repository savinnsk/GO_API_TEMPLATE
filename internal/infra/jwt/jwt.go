package infra

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/savinnsk/api-template-go/configs"
	"github.com/savinnsk/api-template-go/internal/domain"

)

func GenerateJwt(userEmail string)(*domain.LoginResponse,error){
	env,err := configs.LoadEnv()
	
	if err != nil {
		panic(err)
	}

	claims := jwt.MapClaims{
		"sub":userEmail,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}


	signature := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	
	token ,err := signature.SignedString([]byte(env.JwtSecret))

	if err != nil {
		fmt.Errorf("%s", err.Error())
		return nil,err
	}

	if err != nil {
		fmt.Errorf("%s", err.Error())
		return nil,err
	}
	return  &domain.LoginResponse{Token: token}, nil

}