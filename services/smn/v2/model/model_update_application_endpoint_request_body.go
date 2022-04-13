package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateApplicationEndpointRequestBody struct {
	// 设备是否可用，值为true或false字符串。

	Enabled *string `json:"enabled,omitempty"`
	// 用户自定义数据，最大长度支持UTF-8编码后2048字节。

	UserData *string `json:"user_data,omitempty"`
}

func (o UpdateApplicationEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateApplicationEndpointRequestBody", string(data)}, " ")
}
