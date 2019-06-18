package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"../service"
	"../util"
)

var contactService service.ContactService

//

func Addfriend(w http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	from, err1 := strconv.ParseInt(request.PostForm.Get("fromid"), 10, 64)
	to, err2 := strconv.ParseInt(request.PostForm.Get("toid"), 10, 64)

	if err1 != nil || err2 != nil {
		util.Resp(w, 1003, nil, err1.Error())
		return
	}

	//调用service
	err := contactService.AddFriend(from, to)
	//
	if err != nil {
		util.Resp(w, 1007, nil, err.Error())
	} else {
		util.Resp(w, 0, nil, "succ")
	}
}

func LoadFriend(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	userid, _ := strconv.ParseInt(request.PostForm.Get("userid"), 10, 64)
	fmt.Println(userid)
	users := contactService.SearchFriend(userid)
	util.Resp(w, 0, users, "成功")
}

func JoinCommunity(w http.ResponseWriter, request *http.Request) {

	userid, _ := strconv.ParseInt(request.PostForm.Get("userid"), 10, 64)
	dstid, _ := strconv.ParseInt(request.PostForm.Get("dstid"), 10, 64)

	err := contactService.JoinCommunity(userid, dstid)
	//todo 刷新用户的群组信息
	AddGroupId(userid, dstid)
	if err != nil {
		util.Resp(w, 0, nil, err.Error())
	} else {
		util.Resp(w, 0, nil, "")
	}
}
