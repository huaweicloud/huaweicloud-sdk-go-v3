package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CancelShareNewRequestBody struct {

	// 共享链接ID
	SharedConnId *string `json:"shared_conn_id,omitempty"`

	// 用户列表
	Users *[]CancelShareNewRequestBodyUsers `json:"users,omitempty"`
}

func (o CancelShareNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelShareNewRequestBody struct{}"
	}

	return strings.Join([]string{"CancelShareNewRequestBody", string(data)}, " ")
}
