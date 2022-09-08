package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestartEncodeServerRequestBody struct {

	// 待重启的编码服务的ID
	EncodeServerIds []string `json:"encode_server_ids"`
}

func (o RestartEncodeServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartEncodeServerRequestBody struct{}"
	}

	return strings.Join([]string{"RestartEncodeServerRequestBody", string(data)}, " ")
}
