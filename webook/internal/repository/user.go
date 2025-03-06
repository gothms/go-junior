package repository

import (
	"github.com/gin-gonic/gin"
	"go-junior/webook/internal/domain"
	"go-junior/webook/internal/repository/dao"
	"time"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	// ErrUserNotFound repository 是与业务强相关的，返回错误应告诉什么业务报错
	ErrUserNotFound = dao.ErrRecordNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

func (repo *UserRepository) Create(ctx *gin.Context, u domain.User) error {
	return repo.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (repo *UserRepository) FindByEmail(ctx *gin.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}
func (repo *UserRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}
}
func (repo *UserRepository) toDomainWithPersonalInfo(u dao.User) domain.User {
	return domain.User{
		Nickname: u.Nickname,
		Birthday: time.Unix(u.Birthday, 0),
		Personal: u.Personal,
	}
}

func (repo *UserRepository) UpdateNonZeroFields(ctx *gin.Context, id int64, nickname string, birthday int64, personal string) error {
	return repo.dao.UpdateNonZeroFields(ctx, id, nickname, birthday, personal)
}

func (repo *UserRepository) FindUserById(ctx *gin.Context, id int64) (domain.User, error) {
	u, err := repo.dao.FindUserById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomainWithPersonalInfo(u), nil
}
