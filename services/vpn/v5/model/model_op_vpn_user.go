package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpVpnUser struct {

	// 用户ID
	Id string `json:"id"`
}

func (o OpVpnUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpVpnUser struct{}"
	}

	return strings.Join([]string{"OpVpnUser", string(data)}, " ")
}
