package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSharedInfoNewRequestBody 更新共享信息请求体
type UpdateSharedInfoNewRequestBody struct {

	// 修改后共享的新用户ID
	UserId string `json:"user_id"`

	// 修改后共享的新用户名
	NewUserName string `json:"new_user_name"`
}

func (o UpdateSharedInfoNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSharedInfoNewRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSharedInfoNewRequestBody", string(data)}, " ")
}
