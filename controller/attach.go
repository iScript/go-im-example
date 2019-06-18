package controller

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"../util"
)

func init() {
	os.MkdirAll("./mnt", os.ModePerm)
}
func Upload(w http.ResponseWriter, r *http.Request) {
	UploadLocal(w, r)

}

func UploadLocal(writer http.ResponseWriter,
	request *http.Request) {
	//todo 获得上传的源文件s
	srcfile, head, err := request.FormFile("file")
	if err != nil {
		util.Resp(writer, 1008, nil, "upload fail")
	}

	//todo 创建一个新文件d
	suffix := ".png"
	//如果前端文件名称包含后缀 xx.xx.png
	ofilename := head.Filename
	tmp := strings.Split(ofilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	//如果前端指定filetype
	//formdata.append("filetype",".png")
	filetype := request.FormValue("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}
	//time.Now().Unix()
	filename := fmt.Sprintf("%d%04d%s",
		time.Now().Unix(), rand.Int31(),
		suffix)
	dstfile, err := os.Create("./mnt/" + filename)
	if err != nil {
		util.Resp(writer, 1009, nil, err.Error())
		return
	}

	//todo 将源文件内容copy到新文件
	_, err = io.Copy(dstfile, srcfile)
	if err != nil {
		util.Resp(writer, 1009, nil, err.Error())
		return
	}
	//todo 将新文件路径转换成url地址

	url := "/mnt/" + filename
	//todo 响应到前端
	util.Resp(writer, 0, url, "succ")
}
