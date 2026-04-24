package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeUserGroupInfo 查询订阅用户组信息
type SubscribeUserGroupInfo struct {

	// 用户组ID。
	GroupId *string `json:"group_id,omitempty"`

	// 用户组名称。
	GroupName *string `json:"group_name,omitempty"`

	// ai 功能是否启用。 * true： 启用 * false： 不启用
	AiFunc *bool `json:"ai_func,omitempty"`
}

func (o SubscribeUserGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeUserGroupInfo struct{}"
	}

	return strings.Join([]string{"SubscribeUserGroupInfo", string(data)}, " ")
}
