package main

import (
	"fmt"
	"net/http"

	"./controller"
)

func main() {

	fmt.Println("start")

	http.HandleFunc("/login", controller.UserLoginHandle)
	http.HandleFunc("/register", controller.UserRegisterHandle)
	// http.HandleFunc("/contact/loadcommunity", controller.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", controller.LoadFriend)
	// http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)
	http.HandleFunc("/contact/addfriend", controller.Addfriend)
	http.HandleFunc("/chat", controller.Chat) //websocket
	http.HandleFunc("/attach/upload", controller.Upload)
	//前端静态文件
	http.Handle("/", http.FileServer(http.Dir("frontend")))
	http.Handle("/mnt/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)

	//init()
}
