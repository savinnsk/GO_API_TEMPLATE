package repositories

import (
	"github.com/google/uuid"
	"github.com/savinnsk/api-template-go/configs"
	"github.com/savinnsk/api-template-go/internal/domain"
	"github.com/savinnsk/api-template-go/internal/infra/sqlc"

)
func CreateUser(userDto domain.UserDto ) (sqlc.User ,error) {
	ctx := configs.GetContext();
	db := configs.GetInstanceDB()

	user := sqlc.CreateUserParams{
		ID: uuid.New().String(),
		Name: userDto.Name,
		Password: userDto.Password,
		Email: userDto.Email,
	}

	if err  := db.CreateUser(ctx, user); err != nil {
		return sqlc.User{} , err
	};

	result,err := db.GetUser(ctx,user.ID)
	if (err != nil) {
		return sqlc.User{} , err
	}

	return result,nil
}