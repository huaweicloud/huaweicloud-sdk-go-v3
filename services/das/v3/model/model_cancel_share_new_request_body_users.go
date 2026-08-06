package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CancelShareNewRequestBodyUsers struct {

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`
}

func (o CancelShareNewRequestBodyUsers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelShareNewRequestBodyUsers struct{}"
	}

	return strings.Join([]string{"CancelShareNewRequestBodyUsers", string(data)}, " ")
}
