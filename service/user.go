package service

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"openapi/common"
	"openapi/constant"
	"openapi/global"
	"openapi/modules/user"
	"strconv"
	"strings"
)

// UserService 用户服务
type UserService struct {
}

// GetUserService 获取用户服务
func GetUserService() *UserService {
	return &UserService{}
}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo(username string) *user.UserInfo {
	val, err := global.Rdb.HGet(ctx, constant.RdbUserKey, username).Result()
	if err != nil {
		fmt.Println("获取用户信息异常！", err)
		return nil
	}
	var userInfo user.UserInfo
	err = json.Unmarshal([]byte(val), &userInfo)
	if err != nil {
		fmt.Println("读取用户信息失败！", err)
		return nil
	}
	userInfo.Password = ""
	return &userInfo
}

// UserLogin 用户登录
func (s *UserService) UserLogin(u *user.LoginParam) *user.UserToken {
	val, err := global.Rdb.HGet(ctx, constant.RdbUserKey, u.UserName).Result()
	if err != nil {
		fmt.Println("获取用户信息异常！", err)
		return nil
	}
	var userInfo user.UserInfo
	err = json.Unmarshal([]byte(val), &userInfo)
	if err != nil {
		fmt.Println("读取用户信息失败！", err)
		return nil
	}
	hash := sha1.New()
	data := fmt.Sprintf("%s-%s", u.Password, constant.Salt)
	io.WriteString(hash, data)
	sha1Sam := hash.Sum(nil)
	if fmt.Sprintf("%x", sha1Sam) == userInfo.Password {
		return getToken(&userInfo)
	}
	return nil
}

// UserRegister 用户注册
func (s *UserService) UserRegister(u *user.SignParam) *user.UserToken {
	exists, err := global.Rdb.HExists(ctx, constant.RdbUserKey, u.UserName).Result()
	if err != nil {
		fmt.Println("检查用户信息失败！", err)
		return nil
	}
	if exists {
		fmt.Println("用户名已经注册了!")
		return nil
	}
	hash := sha1.New()
	data := fmt.Sprintf("%s-%s", u.Password, constant.Salt)
	io.WriteString(hash, data)
	sha1Sam := hash.Sum(nil)
	userInfo := func(arg1 string, arg2 string) *user.UserInfo {
		var stringBuilder strings.Builder
		stringBuilder.WriteString(strconv.Itoa(rand.Intn(9) + 1))
		for i := 0; i < 6; i++ {
			stringBuilder.WriteString(strconv.Itoa(rand.Intn(10)))
		}
		userID, err := strconv.Atoi(stringBuilder.String())
		if err != nil {
			return nil
		}
		return &user.UserInfo{
			UserName: arg1,
			Password: arg2,
			UserID:   uint64(userID),
		}
	}(u.UserName, fmt.Sprintf("%x", sha1Sam))
	if userInfo == nil {
		return nil
	}
	val, err := json.Marshal(userInfo)
	isSuccess, err := global.Rdb.HSetNX(ctx, constant.RdbUserKey, u.UserName, val).Result()
	if err != nil || !isSuccess {
		fmt.Println("存储用户信息异常！", userInfo, err)
		return nil
	}
	return getToken(userInfo)
}

func getToken(userInfo *user.UserInfo) *user.UserToken {
	token := common.GeneratorToken(userInfo.UserName)
	return &user.UserToken{
		UserID: userInfo.UserID,
		Token:  token,
	}
}
