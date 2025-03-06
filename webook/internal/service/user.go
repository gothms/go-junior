package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-junior/webook/internal/domain"
	"go-junior/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail = repository.ErrDuplicateEmail
	//ErrUserNotFound          = repository.ErrUserNotFound

	ErrInvalidUserOrPassword = errors.New("账号或密码不对")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (svc *UserService) Signup(ctx *gin.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx *gin.Context, email string, password string) (domain.User, error) {
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	// 检查密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *UserService) UpdateNonSensitiveInfo(ctx *gin.Context, id int64, nickname string, birthday int64, personal string) error {
	return svc.repo.UpdateNonZeroFields(ctx, id, nickname, birthday, personal)
}
func (svc *UserService) FindUserById(ctx *gin.Context, id int64) (domain.User, error) {
	return svc.repo.FindUserById(ctx, id)
}
