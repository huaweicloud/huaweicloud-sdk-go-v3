package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserLastLogin IAM用户最后登录时间。
type UserLastLogin struct {

	// IAM用户最后登录时间。若不存在则表示从未登录过。
	LastLoginAt *sdktime.SdkTime `json:"last_login_at,omitempty"`
}

func (o UserLastLogin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserLastLogin struct{}"
	}

	return strings.Join([]string{"UserLastLogin", string(data)}, " ")
}
