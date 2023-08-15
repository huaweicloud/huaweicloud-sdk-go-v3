package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateServerNameRequestBody struct {

	// 云手机服务器名称，必须为小写字母（a-z）、大写字母（A-Z）、数字（0-9）、中文字符、中划线-、下划线_，且不得超过60个字符。
	ServerName string `json:"server_name"`
}

func (o UpdateServerNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateServerNameRequestBody", string(data)}, " ")
}
