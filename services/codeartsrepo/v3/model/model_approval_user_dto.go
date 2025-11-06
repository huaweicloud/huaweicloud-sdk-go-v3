package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApprovalUserDto struct {

	// id
	Id *int32 `json:"id,omitempty"`

	// 用户名
	Username *string `json:"username,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 昵称
	NickName *string `json:"nick_name,omitempty"`

	// 检视状态
	State *string `json:"state,omitempty"`
}

func (o ApprovalUserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalUserDto struct{}"
	}

	return strings.Join([]string{"ApprovalUserDto", string(data)}, " ")
}
