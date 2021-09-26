package model

import (
	"encoding/json"

	"strings"
)

// 创建应用请求体
type AppReq struct {
	// app名称，仅支持填入utf8格式内容

	AppName string `json:"app_name"`
}

func (o AppReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AppReq struct{}"
	}

	return strings.Join([]string{"AppReq", string(data)}, " ")
}