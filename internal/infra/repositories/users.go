package repositories

import (
	"github.com/savinnsk/api-template-go/configs"
	"github.com/savinnsk/api-template-go/internal/infra/sqlc"

)
func CreateUser(db configs.DbInstance, userParam sqlc.CreateUserParams) error {
	ctx := configs.GetContext();


	err := db.CreateUser(ctx, userParam);

	if err != nil {
		return err
	}
	return nil
}