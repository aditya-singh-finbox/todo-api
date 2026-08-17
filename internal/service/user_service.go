package service

import (
	"fmt"

	"github.com/aditya-singh-finbox/todo-api/internal/model"
	"github.com/aditya-singh-finbox/todo-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(user *model.User) error {
	existingUser, err := s.userRepo.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = string(hashedPassword)

	return s.userRepo.CreateUser(user)
}
func (s *UserService) Login(email, password string) (*model.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	fmt.Printf("Entered Password: '%s'\n", password)
	fmt.Printf("Length: %d\n", len(password))
	if err != nil {
		fmt.Println("User not found:", err)
		return nil, fmt.Errorf("invalid email or password first")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Password comparison failed:", err)
		return nil, fmt.Errorf("invalid email or password second")
	}

	return user, nil
}

func (s *UserService) GetByID(id uint) (*model.User, error) {
	return s.userRepo.GetByID(id)
}
