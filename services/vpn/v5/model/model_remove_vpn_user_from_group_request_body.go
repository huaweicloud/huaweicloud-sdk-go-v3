package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveVpnUserFromGroupRequestBody struct {

	// 用户列表信息
	Users []OpVpnUser `json:"users"`
}

func (o RemoveVpnUserFromGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveVpnUserFromGroupRequestBody struct{}"
	}

	return strings.Join([]string{"RemoveVpnUserFromGroupRequestBody", string(data)}, " ")
}
