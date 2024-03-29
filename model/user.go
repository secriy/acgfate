package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// EncryptCost 加密难度
const EncryptCost = 12

// User 用户账号模型
type User struct {
	UID       int64     `db:"uid"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Nickname  string    `db:"nickname"`
	Email     string    `db:"email"`
	Avatar    string    `db:"avatar"`
	State     uint8     `db:"state"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// UserInfo 用户个人信息模型
type UserInfo struct {
	UID       int64     `db:"uid"`
	Gender    uint8     `db:"gender"`
	Sign      string    `db:"sign"`
	Birthday  time.Time `db:"birthday"`
	Province  string    `db:"province"`
	City      string    `db:"city"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// CurrentUser 获取当前用户模型
func CurrentUser(c *gin.Context) *User {
	ret, ok := c.Get("user")
	if !ok {
		return nil
	}
	user, ok := ret.(*User)
	return user
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), EncryptCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
