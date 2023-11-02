package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertUserInfo struct {

	// 用户uid
	UserId *string `json:"user_id,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`
}

func (o AlertUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertUserInfo struct{}"
	}

	return strings.Join([]string{"AlertUserInfo", string(data)}, " ")
}
