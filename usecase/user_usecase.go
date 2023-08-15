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
	IsAvailableEmail(input *model.CheckEmailAvailable) (bool, error)
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
	register.Role = "user"
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
func (u *userUsecaseImpl) IsAvailableEmail(input *model.CheckEmailAvailable) (bool, error) {
	email := input.Email

	// mencari email yang di input ketika membuat akun
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	// jika tidak ada maka bisa untuk input 
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (u *userUsecaseImpl) UpdateAvatar(id int, user *model.UserModel) error {
	return u.userRepo.UpdateAvatar(id, user)
}

func NewUserUsecase(repo repo.UserRepo) Userusecase{
	return &userUsecaseImpl{
		userRepo: repo,
	}
}