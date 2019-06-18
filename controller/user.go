package controller

import (
	"fmt"
	"math/rand"
	"net/http"

	"../model"
	"../service"
	"../util"
)

var userService service.UserService

func UserRegisterHandle(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	password := request.PostForm.Get("password")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_MEN

	user, err := userService.Register(mobile, password, nickname, avatar, sex)

	if err != nil {
		util.Resp(writer, -1, nil, err.Error())
	} else {
		util.Resp(writer, 0, user, "success")
	}

}

func UserLoginHandle(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("login")

	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	password := request.PostForm.Get("password")

	user, err := userService.Login(mobile, password)

	if err != nil {
		util.Resp(writer, 1006, nil, err.Error())
	} else {
		util.Resp(writer, 0, user, "登录成功")
	}

}
