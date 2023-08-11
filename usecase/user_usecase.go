package usecase

import (
	"fmt"
	"startup/model"
	"startup/repo"

	"golang.org/x/crypto/bcrypt"
)

type Userusecase interface {
	RegisterUser(*model.RegisterUserInput) error
	LoginUser(*model.LoginUser) error
	IsAvalibleEmail(input *model.CheckEmailAvalible) (bool, error)
}

type userUsecaseImpl struct {
	userRepo repo.UserRepo
}


// Registrasi
func (u *userUsecaseImpl) RegisterUser(register *model.RegisterUserInput) error{
	user := model.UserModel{}
	user.Name = register.Name
	user.Email = register.Email
	user.Occupation = register.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.MinCost)
	if err != nil{
		return fmt.Errorf("error an userUsecaseImpl.RegisterUser: %w", err)
	}
	user.PasswordHash = string(passwordHash)
	return u.userRepo.RegisterUser(&user)
}

// Login 
func (u *userUsecaseImpl) LoginUser(input *model.LoginUser) error {
	email := input.Email
	password := input.Password

	// mencari email
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil
	}

	// validasi email ada atau tidak
	if user.ID == 0 {
		return nil
	}

	// mencocokan password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil
	}

	return nil
}

// untuk mengecek apakah email yang didaftarkan sudah ada apa belum
func (u *userUsecaseImpl) IsAvalibleEmail(input *model.CheckEmailAvalible) (bool, error) {
	email := input.Email

	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return false, err
	}
	
	
	return true, nil
}


func NewUserUsecase(repo repo.UserRepo) Userusecase{
	return &userUsecaseImpl{
		userRepo: repo,
	}
}