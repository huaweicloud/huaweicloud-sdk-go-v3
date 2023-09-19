package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestartCloudPhoneServerRequestBody struct {

	// 云手机服务器id列表。
	ServerIds []string `json:"server_ids"`
}

func (o RestartCloudPhoneServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartCloudPhoneServerRequestBody struct{}"
	}

	return strings.Join([]string{"RestartCloudPhoneServerRequestBody", string(data)}, " ")
}
