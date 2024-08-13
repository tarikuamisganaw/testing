package usecases

import (
	"clean-architecture/domain"
	"clean-architecture/infrastructure"
	"clean-architecture/repositories"
	"errors"
)

type UserUsecase interface {
	Register(user domain.User) (domain.User, error)
	Login(username, password string) (string, error)
	GetUsers() ([]domain.User, error)
}

type userUsecase struct {
	userRepo    repositories.UserRepository
	passwordSvc infrastructure.PasswordService
	jwtSvc      infrastructure.JWTService
}

func NewUserUsecase(ur repositories.UserRepository, ps infrastructure.PasswordService, js infrastructure.JWTService) UserUsecase {
	return &userUsecase{
		userRepo:    ur,
		passwordSvc: ps,
		jwtSvc:      js,
	}
}

// GetUsers retrieves all users from the repository
func (uc *userUsecase) GetUsers() ([]domain.User, error) {
	return uc.userRepo.GetUsers()
}

// Register registers a new user
func (u *userUsecase) Register(user domain.User) (domain.User, error) {
	// Hash the user's password before storing it
	hashedPassword, err := u.passwordSvc.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}
	user.Password = hashedPassword

	// Save the user to the repository
	createdUser, err := u.userRepo.Register(user)
	if err != nil {
		return domain.User{}, err
	}
	return createdUser, nil
}

// Login authenticates a user and returns a JWT token if successful
func (u *userUsecase) Login(username, password string) (string, error) {
	// Retrieve the user by username
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Compare the provided password with the stored hashed password
	if err := u.passwordSvc.CheckPasswordHash(user.Password, password); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate a JWT token for the authenticated user
	token, err := u.jwtSvc.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
