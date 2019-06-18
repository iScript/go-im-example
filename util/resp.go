package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` //空interface，可以存储任意类型的数据,omitempty指定该字段为空则不显示
}

func Resp(writer http.ResponseWriter, code int, data interface{}, message string) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)

	h := H{
		Code:    code,
		Message: message,
		Data:    data,
	}
	// fmt.Println("========")
	// fmt.Println(h)

	ret, err := json.Marshal(h)

	if err != nil {
		log.Println(err.Error())
	}

	writer.Write(ret)
}
