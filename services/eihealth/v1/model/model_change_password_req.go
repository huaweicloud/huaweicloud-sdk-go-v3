package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改密码请求体
type ChangePasswordReq struct {

	// 原始密码
	OriginalPassword string `json:"original_password"`

	// 新密码
	NewPassword string `json:"new_password"`

	// 预验证凭证
	Ticket *string `json:"ticket,omitempty"`
}

func (o ChangePasswordReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordReq struct{}"
	}

	return strings.Join([]string{"ChangePasswordReq", string(data)}, " ")
}
