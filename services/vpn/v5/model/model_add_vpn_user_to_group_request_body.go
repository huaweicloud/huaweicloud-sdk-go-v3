package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddVpnUserToGroupRequestBody struct {

	// 用户列表信息
	Users []OpVpnUser `json:"users"`
}

func (o AddVpnUserToGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVpnUserToGroupRequestBody struct{}"
	}

	return strings.Join([]string{"AddVpnUserToGroupRequestBody", string(data)}, " ")
}
