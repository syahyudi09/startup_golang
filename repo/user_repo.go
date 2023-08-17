package repo

import (
	"errors"
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
	UpdateAvatar(model.UserModel) error
}

type userRepoImpl struct {
	db *sql.DB
}

func (u *userRepoImpl) RegisterUser(user *model.UserModel) error{
	query := utils.REGISTER_USER
	
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Role = "user"
	_, err := u.db.Exec(query, user.Name, user.Occupation, user.Email, user.PasswordHash, user.Role , user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("err on userRepoImpl.RegisterUser: %v", err)
	}
	return nil
}

func (u *userRepoImpl) FindByEmail(email string) (model.UserModel, error) {
	query := utils.FIND_BY_EMAIL
	row := u.db.QueryRow(query, email)

	var user model.UserModel
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)
	if err != nil {
		return user, fmt.Errorf("error fetching user by email: %w", err)
	}

	return user, nil
}

func (u *userRepoImpl) GetUserByID(id int) (model.UserModel, error) {
	query := utils.GET_USER_BY_ID
	
	var user model.UserModel
	err := u.db.QueryRow(query, id).Scan(&user.ID)
	if err != nil {
		return user, fmt.Errorf("error an userRepoImpl.GetUserById %d", err)
	}

	if user.ID == 0{
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (u *userRepoImpl) UpdateAvatar(user model.UserModel) error {
	fmt.Println("Updating avatar for user ID:", user.ID)
	query := "UPDATE users SET avatar_filename = $1 WHERE id = $2"
	_, err := u.db.Exec(query, user.AvatarFileName, user.ID)
	if err != nil {
		return fmt.Errorf("error in userRepoImpl.UpdateAvatar: %w", err)
	}
	return nil
}

func NewUserRepo(db *sql.DB) UserRepo{
	return &userRepoImpl{
		db: db,
	}
}