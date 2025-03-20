package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersV5Request Request Object
type ListUsersV5Request struct {

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记，长度为4到400个字符，只包含字母、数字、\"+\"、\"/\"、\"=\"、\"-\"和\"_\"的字符串。
	Marker *string `json:"marker,omitempty"`

	// 用户组ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	GroupId *string `json:"group_id,omitempty"`
}

func (o ListUsersV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersV5Request struct{}"
	}

	return strings.Join([]string{"ListUsersV5Request", string(data)}, " ")
}
