package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunUserRequest 启动用户、启动用户组设置
type RunUserRequest struct {

	// 容器启动用户的user id
	Uid *int32 `json:"uid,omitempty"`

	// 容器启动用户的group id
	Gid *int32 `json:"gid,omitempty"`
}

func (o RunUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunUserRequest struct{}"
	}

	return strings.Join([]string{"RunUserRequest", string(data)}, " ")
}
