package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunUserInfo 启动用户、启动用户组设置
type RunUserInfo struct {

	// 容器启动用户的user id
	Uid *int32 `json:"uid,omitempty"`

	// 容器启动用户的group id
	Gid *int32 `json:"gid,omitempty"`

	// 容器启动用户的user name
	UserName *string `json:"user_name,omitempty"`

	// 容器启动用户的group name
	GroupName *string `json:"group_name,omitempty"`
}

func (o RunUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunUserInfo struct{}"
	}

	return strings.Join([]string{"RunUserInfo", string(data)}, " ")
}
