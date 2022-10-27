package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApprovalMergeRequestApproversItem struct {

	// 用户id
	Id *float64 `json:"id,omitempty"`

	// 用户名
	Name *string `json:"name,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	// 用户状态
	State *string `json:"state,omitempty"`

	// 用户iamId
	Username *string `json:"username,omitempty"`
}

func (o ApprovalMergeRequestApproversItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalMergeRequestApproversItem struct{}"
	}

	return strings.Join([]string{"ApprovalMergeRequestApproversItem", string(data)}, " ")
}
