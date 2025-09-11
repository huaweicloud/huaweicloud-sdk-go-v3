package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MountUser 挂载用户信息。
type MountUser struct {

	// 用户ID(-1~65534的非0整数)
	UserId int32 `json:"user_id"`

	// 用户组ID(-1~65534的非0整数)
	UserGroupId int32 `json:"user_group_id"`
}

func (o MountUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MountUser struct{}"
	}

	return strings.Join([]string{"MountUser", string(data)}, " ")
}
