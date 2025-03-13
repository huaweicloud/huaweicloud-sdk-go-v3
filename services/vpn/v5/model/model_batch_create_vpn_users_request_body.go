package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateVpnUsersRequestBody struct {

	// 用户列表
	Users []CreateVpnUser `json:"users"`
}

func (o BatchCreateVpnUsersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVpnUsersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateVpnUsersRequestBody", string(data)}, " ")
}
