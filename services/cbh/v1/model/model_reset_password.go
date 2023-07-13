package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetPassword 云堡垒机实例修改admin用户密码
type ResetPassword struct {

	// admin用户修改后的新密码，8-32位，大写字母、小写字母、数字和特殊字符。
	NewPassword string `json:"new_password"`

	// 云堡垒机实例ID，使用UUID格式。
	ServerId string `json:"server_id"`
}

func (o ResetPassword) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPassword struct{}"
	}

	return strings.Join([]string{"ResetPassword", string(data)}, " ")
}
