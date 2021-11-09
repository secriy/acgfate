package service

import (
	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterService struct {
	Username string `json:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" binding:"required,ascii,min=8,max=16"`
	Nickname string `json:"nickname" binding:"required,min=2,max=15"`
	Email    string `json:"email" binding:"required,email,min=3,max=100"`
}

// Register 用户注册服务
func (service *RegisterService) Register() sz.Response {
	var user model.User
	dao := new(database.UserDao)
	// 判断用户名是否被占用
	if _, err := dao.QueryByUname(service.Username); err == nil {
		return sz.CodeResponse(sz.CodeRegNameExist)
	}
	// 判断邮箱是否被占用
	if _, err := dao.QueryByEmail(service.Email); err == nil {
		return sz.CodeResponse(sz.CodeEmailExist)
	}
	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		zap.S().Errorf("%s: %s", sz.Msg(sz.CodePasswdEncryptErr), err)
		return sz.ErrorResponse()
	}
	// 创建用户账号记录
	user.Username = service.Username
	user.Nickname = service.Nickname
	user.Email = service.Email
	if err := dao.InsertRow(user); err != nil {
		zap.S().Errorf("创建用户失败: %s", err)
		return sz.ErrorResponse()
	}
	return sz.SuccessResponse()
}
