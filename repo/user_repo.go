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
	FindByEmail(string) (model.UserModel, error)
	GetUserByID(int) (model.UserModel, error)
	UpdateAvatar(int, *model.UserModel) error
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

func (u *userRepoImpl) FindByEmail(email string) (model.UserModel, error) {
	query := utils.FIND_BY_EMAIL
	row := u.db.QueryRow(query, email)

	var user model.UserModel
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return user, fmt.Errorf("error fetching user by email: %w", err)
	}

	return user, nil
}

func (u *userRepoImpl) GetUserByID(id int) (model.UserModel, error) {
	query := utils.GET_USER_BY_ID
	row := u.db.QueryRow(query, id)

	var user model.UserModel
	err := row.Scan(user.ID)
	if err != nil {
		return user, fmt.Errorf("error an userRepoImpl.GetUserById %d", err)
	}

	return user, nil
}

func (u *userRepoImpl) UpdateAvatar(id int, cust *model.UserModel) error {
	_, err := u.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("error an userRepoImpl.UpdateAvatar %w", err)
	}

	query := "UPDATE users SET "
	_, err = u.db.Exec(query, &cust.AvatarFileName, &cust.ID)
	if err != nil {
		return fmt.Errorf("error an userRepoImpl.UpdateAvatar %w", err)
	}
	return nil
}
func NewUserRepo(db *sql.DB) UserRepo{
	return &userRepoImpl{
		db: db,
	}
}