package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 挂载用户信息。
type MountUser struct {

	// 用户ID(-1~65534的非0整数)
	UserId string `json:"user_id" xml:"user_id"`

	// 用户组ID(-1~65534的非0整数)
	UserGroupId string `json:"user_group_id" xml:"user_group_id"`
}

func (o MountUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MountUser struct{}"
	}

	return strings.Join([]string{"MountUser", string(data)}, " ")
}
