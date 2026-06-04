package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeUserBasicInfo 查询订阅用户信息
type SubscribeUserBasicInfo struct {

	// 用户id。
	UserId *string `json:"user_id,omitempty"`

	// 桌面用户名。
	UserName *string `json:"user_name,omitempty"`

	// 用户手机号。
	UserPhone *string `json:"user_phone,omitempty"`

	// ai 功能是否启用。 * true： 启用 * false： 不启用
	AiFunc *bool `json:"ai_func,omitempty"`
}

func (o SubscribeUserBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeUserBasicInfo struct{}"
	}

	return strings.Join([]string{"SubscribeUserBasicInfo", string(data)}, " ")
}
