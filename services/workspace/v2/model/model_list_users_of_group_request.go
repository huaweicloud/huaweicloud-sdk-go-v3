package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersOfGroupRequest Request Object
type ListUsersOfGroupRequest struct {

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 用户组ID。
	GroupId string `json:"group_id"`

	// 用于分页查询，返回桌面数量限制。如果不指定或为0，默认1000，最大1000。
	Limit *string `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *string `json:"offset,omitempty"`
}

func (o ListUsersOfGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersOfGroupRequest struct{}"
	}

	return strings.Join([]string{"ListUsersOfGroupRequest", string(data)}, " ")
}
