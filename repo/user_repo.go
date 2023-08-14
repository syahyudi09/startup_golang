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

func (u *userRepoImpl) FindByEmail(email string) (*model.UserModel, error) {
	query := "SELECT id, name, email FROM users WHERE email = $1"
	row := u.db.QueryRow(query, email)

	var user model.UserModel
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("error fetching user by email: %w", err)
	}

	return &user, nil
}

func NewUserRepo(db *sql.DB) UserRepo{
	return &userRepoImpl{
		db: db,
	}
}