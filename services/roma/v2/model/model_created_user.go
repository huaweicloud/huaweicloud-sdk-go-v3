package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建时的用户
type CreatedUser struct {
	// 用户ID

	UserId *string `json:"user_id,omitempty"`
	// 用户名

	UserName *string `json:"user_name,omitempty"`
}

func (o CreatedUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatedUser struct{}"
	}

	return strings.Join([]string{"CreatedUser", string(data)}, " ")
}
