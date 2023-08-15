package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetInstancePasswordBody struct {

	// 重置的新密码
	NewPassword *string `json:"new_password,omitempty"`

	// 是否重置为免密码访问缓存实例
	NoPasswordAccess *bool `json:"no_password_access,omitempty"`
}

func (o ResetInstancePasswordBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetInstancePasswordBody struct{}"
	}

	return strings.Join([]string{"ResetInstancePasswordBody", string(data)}, " ")
}
