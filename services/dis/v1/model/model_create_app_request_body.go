package model

import (
	"encoding/json"

	"strings"
)

type CreateAppRequestBody struct {
	// APP的名称，用户数据消费程序的唯一标识符。  应用名称由字母、数字、下划线和中划线组成，长度为1～200个字符。

	AppName string `json:"app_name"`
}

func (o CreateAppRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAppRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAppRequestBody", string(data)}, " ")
}
