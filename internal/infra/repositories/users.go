package repositories

import (
	"github.com/savinnsk/api-template-go/internal/infra/sqlc"
	"github.com/savinnsk/api-template-go/configs"

)
func CreateUser(userParam sqlc.CreateUserParams) (sqlc.User ,error) {
	ctx := configs.GetContext();
	db := configs.GetInstanceDB()

	if err  := db.CreateUser(ctx, userParam); err != nil {
		return sqlc.User{} , err
	};

	user,err := db.GetUser(ctx,userParam.ID)
	if (err != nil) {
		return sqlc.User{} , err
	}

	return user,nil
}