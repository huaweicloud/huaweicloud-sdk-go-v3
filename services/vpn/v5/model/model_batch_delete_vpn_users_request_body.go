package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteVpnUsersRequestBody struct {

	// 用户列表
	Users []OpVpnUser `json:"users"`
}

func (o BatchDeleteVpnUsersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVpnUsersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteVpnUsersRequestBody", string(data)}, " ")
}
