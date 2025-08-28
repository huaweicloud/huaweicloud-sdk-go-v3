package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginProfile IAM用户登录信息。
type LoginProfile struct {

	// IAM用户ID。
	UserId string `json:"user_id"`

	// IAM用户下次登录时是否需要修改密码。
	PasswordResetRequired bool `json:"password_reset_required"`

	// IAM用户密码过期时间。
	PasswordExpiresAt *sdktime.SdkTime `json:"password_expires_at,omitempty"`

	// IAM用户登录信息创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`
}

func (o LoginProfile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginProfile struct{}"
	}

	return strings.Join([]string{"LoginProfile", string(data)}, " ")
}
