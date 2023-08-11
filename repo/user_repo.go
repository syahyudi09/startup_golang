package repo

import (
	"database/sql"
	"fmt"
	"startup/model"
	"startup/utils"
	"time"
)

type UserRepo interface {
	RegisterUser(*model.UserModel) error
	FindByEmail(string) (*model.UserModel, error)
}

type userRepoImpl struct {
	db *sql.DB
}

func (u *userRepoImpl) RegisterUser(user *model.UserModel) error{
	query := utils.REGISTER_USER
	
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := u.db.Exec(query, user.Name, user.Occupation, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("err on userRepoImpl.RegisterUser: %v", err)
	}
	return nil
}

// mencari user berdasarkan email
func (u *userRepoImpl) FindByEmail(email string) (*model.UserModel, error) {

	query := utils.FIND_BY_EMAIL

	var user *model.UserModel
	err := u.db.QueryRow(query, user.Email)
	if err != nil {
		return user, fmt.Errorf("email not found %v", err)
	}
	return user, nil
}

func NewUserRepo(db *sql.DB) UserRepo{
	return &userRepoImpl{
		db: db,
	}
}