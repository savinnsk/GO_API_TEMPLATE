package repositories

import (
	"github.com/google/uuid"
	"github.com/savinnsk/api-template-go/configs"
	"github.com/savinnsk/api-template-go/internal/domain"
	"github.com/savinnsk/api-template-go/internal/infra/sqlc"

)

func CreateUser(userDto domain.UserDto) (sqlc.User, error) {
	ctx := configs.GetContext()
	db := configs.GetInstanceDB()

	user := sqlc.CreateUserParams{
		ID:       uuid.New().String(),
		Name:     userDto.Name,
		Password: userDto.Password,
		Email:    userDto.Email,
	}

	if err := db.CreateUser(ctx, user); err != nil {
		return sqlc.User{}, err
	}

	result, err := db.GetUser(ctx, user.ID)
	if err != nil {
		return sqlc.User{}, err
	}

	return result, nil
}

func GetUserById(id string) (sqlc.User, error) {
	ctx := configs.GetContext()
	db := configs.GetInstanceDB()

	result, err := db.GetUser(ctx, id)
	if err != nil {
		return sqlc.User{}, err
	}
	return result, nil

}

func GetUsersAscPagination(pag sqlc.ListUsersAscParams) ([]sqlc.User, error) {
	ctx := configs.GetContext()
	db := configs.GetInstanceDB()

	resultAsc, err := db.ListUsersAsc(ctx, sqlc.ListUsersAscParams{Limit: pag.Limit, Offset: pag.Offset})
	if err != nil {
		return []sqlc.User{}, err
	}
	return resultAsc, nil

}

func GetUsersDescPagination(pag sqlc.ListUsersDescParams) ([]sqlc.User, error) {
	ctx := configs.GetContext()
	db := configs.GetInstanceDB()
	resultDesc, err := db.ListUsersDesc(ctx, sqlc.ListUsersDescParams{Limit: pag.Limit, Offset: pag.Offset})
	if err != nil {
		return []sqlc.User{}, err
	}
	return resultDesc, nil

}

func GetUsers() ([]sqlc.User, error) {
	ctx := configs.GetContext()
	db := configs.GetInstanceDB()
	result, err := db.ListUsers(ctx)
	if err != nil {
		return []sqlc.User{}, err
	}
	return result, nil

}

func GetByEmail(email string) (sqlc.User, error) {
	ctx := configs.GetContext()
	db := configs.GetInstanceDB()
	result, err := db.GetUserByEmail(ctx,email)
	if err != nil {
		return sqlc.User{}, err
	}
	return result, nil

}

func UpdateUser(userDto sqlc.User) (sqlc.User, error) {
	ctx := configs.GetContext()
	db := configs.GetInstanceDB()

	user := sqlc.UpdateUserParams{
		ID: userDto.ID,
		Name:     userDto.Name,
		Password: userDto.Password,
		Email:    userDto.Email,
	}

	if err := db.UpdateUser(ctx, user); err != nil {
		return sqlc.User{}, err
	}

	result, err := db.GetUser(ctx, user.ID)
	if err != nil {
		return sqlc.User{}, err
	}

	return result, nil
}

func DeleteUser(id string) (*string, error) {
	ctx := configs.GetContext()
	db := configs.GetInstanceDB()
	err := db.DeleteUser(ctx,id)
	if err != nil {
		return nil, err
	}
	
	return nil, nil

}