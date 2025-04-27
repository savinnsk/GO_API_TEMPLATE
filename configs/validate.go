package configs

import "github.com/go-playground/validator/v10"


var Validate *validator.Validate

func ValidateInstance(){
	
	Validate = validator.New(validator.WithRequiredStructEnabled())

}