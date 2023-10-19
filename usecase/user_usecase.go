package usecase

import (
	"errors"
	"fmt"
	"startup/middleware"
	"startup/model"
	"startup/repo"

	"golang.org/x/crypto/bcrypt"

)

type Userusecase interface {
	RegisterUser(*model.RegisterUserInput) error
	LoginUser(model.LoginUser) (string, error)
	IsAvailableEmail(*model.CheckEmailAvailable) (bool, error)
	UpdateAvatar(int, string, *model.UserModel) error
	GetUserByID(int) (model.UserModel, error)
}

type userUsecaseImpl struct {
	userRepo repo.UserRepo
	auth     middleware.Auth
}

func (u *userUsecaseImpl) RegisterUser(register *model.RegisterUserInput) error {
	user := model.UserModel{}
	user.Name = register.Name
	user.Email = register.Email
	user.Occupation = register.Occupation
	user.Role = "user"

	// Hash password sebelum disimpan ke dalam database
	passHash, err := generatePasswordHash(register.Password)
	if err != nil {
		return fmt.Errorf("err  %w", err)
	}
	user.PasswordHash = passHash

	// Simpan pengguna ke dalam database
	return u.userRepo.RegisterUser(&user)
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

func (u *userUsecaseImpl) LoginUser(input model.LoginUser) (string, error) {
	email := input.Email
	password := input.Password

	// Mencari pengguna berdasarkan email
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", fmt.Errorf("failed to find user by email: %w", err)
	}

	// Validasi apakah pengguna ditemukan
	if user.ID == 0 {
		return "", errors.New("user not found")
	}

	fmt.Println("user.ID:", user.ID)
	fmt.Println("Email:", user.Email)

	// Membandingkan password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("invalid password")
	}

	// Menghasilkan token
	token, err := u.auth.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
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

func (u *userUsecaseImpl) GetUserByID(userID int) (model.UserModel, error) {
	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return user, nil
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (u *userUsecaseImpl) UpdateAvatar(id int, fileLocation string, user *model.UserModel) error {
	user.AvatarFileName = fileLocation

	// Simpan foto dengan memanggil UpdateAvatar dengan pointer user yang diperbarui
	err := u.userRepo.UpdateAvatar(id, user)
	if err != nil {
		return err
	}

	return nil
}

func NewUserUsecase(repo repo.UserRepo, auth middleware.Auth) Userusecase {
	return &userUsecaseImpl{
		userRepo: repo,
		auth:     auth,
	}
}
