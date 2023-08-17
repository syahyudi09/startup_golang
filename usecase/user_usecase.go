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
	LoginUser(model.LoginUser)(string, error)
	IsAvailableEmail(*model.CheckEmailAvailable) (bool, error)
	UpdateAvatar(int, string) error
}

type userUsecaseImpl struct {
	userRepo repo.UserRepo
	auth middleware.AuhtMiddleware
}


// Registrasi
// func (u *userUsecaseImpl) RegisterUser(register *model.RegisterUserInput) error{
// 	user := model.UserModel{}
// 	user.Name = register.Name
// 	user.Email = register.Email
// 	user.Occupation = register.Occupation
// 	user.Role = register.Role
// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
// 	if err != nil{
// 		return fmt.Errorf("error an userUsecaseImpl.RegisterUser: %w", err)
// 	}
// 	user.PasswordHash = string(passwordHash)
// 	return u.userRepo.RegisterUser(&user)
// }

func (u *userUsecaseImpl) RegisterUser(register *model.RegisterUserInput) error {
	user := model.UserModel{}
	user.Name = register.Name
	user.Email = register.Email
	user.Occupation = register.Occupation
	user.Role = register.Role

	// Hash password sebelum disimpan ke dalam database
	passHash, err := generatePasswordHash(register.Password)
	if err != nil{
		return fmt.Errorf("err  %w", err)
	}
	user.PasswordHash = passHash
	

	// Simpan pengguna ke dalam database
	return u.userRepo.RegisterUser(&user)
}

func generatePasswordHash(password string) (string , error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return "", nil
	}
	return string(hash), nil
}

// Login 
// func (u *userUsecaseImpl) LoginUser(input model.LoginUser) (string, error) {
// 	email := input.Email
// 	password := input.Password

// 	// Mencari pengguna berdasarkan email
// 	user, err := u.userRepo.FindByEmail(email)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to find user by email: %w", err)
// 	}

// 	// Validasi apakah pengguna ditemukan
// 	if user.ID == 0 {
// 		return "", errors.New("user not found")
// 	}

// 	fmt.Println("user.ID:", user.ID)
// 	fmt.Println("Email:", user.Email)
// 	fmt.Println("input.Password:", input.Password)

// 	// Membandingkan password
// 	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", fmt.Errorf(err.Error())
// 	}

// 	// Menghasilkan token
// 	token, err := u.auth.GenerateToken(user.ID)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to generate token: %w", err)
// 	}

// 	return token, nil
// }

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

func (u *userUsecaseImpl) UpdateAvatar(id int, fileLocation string) error {
		// Mencari user dengan ID tertentu
		user, err := u.userRepo.GetUserByID(id)
		if err != nil {
			return nil
		}
	
		user.AvatarFileName = fileLocation
	
		// Simpan foto dengan memanggil UpdateAvatar dengan pointer user yang diperbarui
		err = u.userRepo.UpdateAvatar(user)
		if err != nil {
			return err
		}
	
		return nil
	}
	

func NewUserUsecase(repo repo.UserRepo, auth middleware.AuhtMiddleware) Userusecase{
	return &userUsecaseImpl{
		userRepo: repo,
		auth: auth,
	}
}