package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTwoFactorLoginConfigRequestBody set two factor login config
type SetTwoFactorLoginConfigRequestBody struct {

	// 是否开启双因子认证
	Enabled bool `json:"enabled"`

	// 认证类型 - sms ：短信邮件验证 - code ：验证码验证
	AuthType string `json:"auth_type"`

	// 主题别名
	TopicDisplayName string `json:"topic_display_name"`

	// 主题的唯一资源标识
	TopicUrn string `json:"topic_urn"`

	// 服务器列表
	HostIdList []string `json:"host_id_list"`
}

func (o SetTwoFactorLoginConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTwoFactorLoginConfigRequestBody struct{}"
	}

	return strings.Join([]string{"SetTwoFactorLoginConfigRequestBody", string(data)}, " ")
}
