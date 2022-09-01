package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户信息
type UserInfo struct {

	// 用户id
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 用户名
	UserName *string `json:"user_name,omitempty" xml:"user_name"`
}

func (o UserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserInfo struct{}"
	}

	return strings.Join([]string{"UserInfo", string(data)}, " ")
}
