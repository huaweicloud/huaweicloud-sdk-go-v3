package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerAppInfoLastUpdatedUser 最后更新用户信息
type ServerAppInfoLastUpdatedUser struct {

	// 最后修改者的用户ID
	UserId *string `json:"user_id,omitempty"`

	// 最后修改者的用户名
	UserName *string `json:"user_name,omitempty"`
}

func (o ServerAppInfoLastUpdatedUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerAppInfoLastUpdatedUser struct{}"
	}

	return strings.Join([]string{"ServerAppInfoLastUpdatedUser", string(data)}, " ")
}
