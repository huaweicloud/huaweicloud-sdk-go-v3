package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CBH服务实例修改密码
type ResetPassword struct {

	// 新密码
	NewPassword string `json:"new_password"`

	// Cbh Server Id
	ServerId string `json:"server_id"`
}

func (o ResetPassword) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPassword struct{}"
	}

	return strings.Join([]string{"ResetPassword", string(data)}, " ")
}
