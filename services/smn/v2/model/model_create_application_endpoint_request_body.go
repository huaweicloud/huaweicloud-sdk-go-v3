package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateApplicationEndpointRequestBody struct {
	// 移动应用设备token，最大长度512个字节。

	Token string `json:"token"`
	// 用户自定义数据，最大长度支持UTF-8编码后2048字节。

	UserData string `json:"user_data"`
}

func (o CreateApplicationEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"CreateApplicationEndpointRequestBody", string(data)}, " ")
}
