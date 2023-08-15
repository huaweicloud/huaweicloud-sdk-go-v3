package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestAssigneeListItem struct {

	// 用户id
	Id *float64 `json:"id,omitempty"`

	// 用户名
	Name *string `json:"name,omitempty"`

	// 用户状态
	State *string `json:"state,omitempty"`

	// 用户iamId
	Username *string `json:"username,omitempty"`
}

func (o MergeRequestAssigneeListItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestAssigneeListItem struct{}"
	}

	return strings.Join([]string{"MergeRequestAssigneeListItem", string(data)}, " ")
}
