package services

import (
	"camvan/models"
	"camvan/repositories"
	"camvan/token"
	"errors"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUpUser(user *models.User) error
	SignInUser(phone, password string) (string, string, error)
	GetAllUser() ([]models.User, error)
	DeleteUser(id uint) error
	RefreshAccessToken(refreshToken string) (string, error)
}

type userService struct {
	userRepo     repositories.UserRepository
	employeeRepo repositories.EmployeeRepository
}

func NewUserService(userRepo repositories.UserRepository, employeeRepo repositories.EmployeeRepository) UserService {
	return &userService{userRepo, employeeRepo}
}

func (s *userService) SignUpUser(user *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	var employeeID = randomIDEm()
	if user.IsEmployee {
		employee := &models.Employee{
			IDEm:            employeeID,
			Phone:           user.Phone,
			Password:        user.Password,
			Name:            "van",
			SubDepartmentID: 2,
		}
		if err := s.employeeRepo.Create(employee); err != nil {
			return err
		}
	}
	user.IDUser = employeeID
	user.ConfirmPass = user.Password
	return s.userRepo.SignUp(user)
}

func (s *userService) SignInUser(phone, password string) (string, string, error) {
	user, err := s.userRepo.SignIn(phone)
	if err != nil {
		return "", "", errors.New("invalid phone or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid phone or password compare")
	}

	if user.IsEmployee {
		_, err := s.employeeRepo.FindByPhonePass(phone, user.Password)
		if err != nil {
			return "", "", errors.New("invalid phone or password")
		}
	}

	accessToken, err := token.GenerateAccessToken(user.Phone, time.Minute*15)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := token.GenerateRefreshToken(user.Phone, time.Hour*24)
	if err != nil {
		return "", "", err
	}
	// user.AccessToken = accessToken
	// user.RefreshToken = refreshToken
	return accessToken, refreshToken, nil
}
func (s *userService) RefreshAccessToken(refreshToken string) (string, error) {

	claims, err := token.ParseJWT(refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	accessToken, err := token.GenerateAccessToken(claims.Phone, time.Minute*15)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *userService) GetAllUser() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func randomIDEm() uint {
	return uint(rand.Intn(1000-1) + 1)
}
