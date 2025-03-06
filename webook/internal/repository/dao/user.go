package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

type User struct {
	// autoIncrement：不存在页分页；适合范围查询；充分利用操作系统预读机制
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string

	Nickname string `gorm:"type=varchar(128)"` // 必须大写，否则不会在数据库中创建字段
	Birthday int64
	Personal string `gorm:"type=varchar(4096)"`

	// 时区：服务器时区、Go应用时区、数据库时区，多个时区之间转换非常容易出错，则整个系统内部直接定义 UTC 0 的毫秒数
	// 则让前端处理时区，或者在返回数据给前端时处理时区
	Ctime int64
	Utime int64
}

func (dao *UserDAO) Insert(ctx *gin.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime, u.Utime = now, now
	//fmt.Printf("%v \n", u)
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mge, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr uint16 = 1062
		if mge.Number == duplicateErr {
			// 邮箱冲突
			return ErrDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx *gin.Context, email string) (User, error) {
	var u User
	return u, dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
}

func (dao *UserDAO) UpdateNonZeroFields(ctx *gin.Context, id int64, nickname string, birthday int64, personal string) error {
	//fmt.Println("edit:", id, nickname, birthday, personal)
	return dao.db.WithContext(ctx).Where("id=?", id).
		Updates(User{Nickname: nickname, Birthday: birthday, Personal: personal}).Error
}
func (dao *UserDAO) FindUserById(ctx *gin.Context, id int64) (User, error) {
	var u User
	return u, dao.db.WithContext(ctx).Where("id=?", id).Find(&u).Error
}
