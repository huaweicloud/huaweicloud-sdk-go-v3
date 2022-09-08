package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建用户信息
type ServerAppInfoCreatedUser struct {

	// 创建应用的用户ID
	UserId *string `json:"user_id,omitempty"`

	// 创建应用的用户名称
	UserName *string `json:"user_name,omitempty"`
}

func (o ServerAppInfoCreatedUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerAppInfoCreatedUser struct{}"
	}

	return strings.Join([]string{"ServerAppInfoCreatedUser", string(data)}, " ")
}
