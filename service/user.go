package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"../model"
	"github.com/iScript/im-go/util"
)

type UserService struct {
}

func (s *UserService) Register(mobile, password, nickname, avatar, sex string) (user model.User, err error) {

	tmp := model.User{}
	_, err = DbEngine.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}

	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}

	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(password, tmp.Salt)
	tmp.Createat = time.Now()
	//token 可以是一个随机数
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())

	//插入 InserOne
	_, err = DbEngine.InsertOne(&tmp)

	return tmp, err

}

func (s *UserService) Login(mobile, password string) (user model.User, err error) {

	//首先通过手机号查询用户
	tmp := model.User{}
	DbEngine.Where("mobile = ?", mobile).Get(&tmp)
	//如果没有找到
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}
	//查询到了比对密码
	if !util.ValidatePasswd(password, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}

	//刷新token,安全
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token

	DbEngine.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp, nil
}

//查找某个用户
func (s *UserService) Find(
	userId int64) (user model.User) {

	//首先通过手机号查询用户
	tmp := model.User{}
	DbEngine.ID(userId).Get(&tmp)
	return tmp
}
