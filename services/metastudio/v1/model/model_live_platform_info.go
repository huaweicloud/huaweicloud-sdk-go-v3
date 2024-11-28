package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LivePlatformInfo 直播平台信息
type LivePlatformInfo struct {

	// 平台ID
	PlatformId *string `json:"platform_id,omitempty"`

	AccessType *AccessTypeEnum `json:"access_type,omitempty"`

	// 直播平台名称
	Name *string `json:"name,omitempty"`

	AuthorizationInfo *PlatformAuthorizationInfo `json:"authorization_info,omitempty"`

	AuthConfig *CustomPlatformAuthConfig `json:"auth_config,omitempty"`

	// 自定义直播平台回调配置。同一种类型仅保留一个配置，如果配置多个会随机保存一个。
	CallbackConfig *[]StandardPlatformApiConfig `json:"callback_config,omitempty"`
}

func (o LivePlatformInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LivePlatformInfo struct{}"
	}

	return strings.Join([]string{"LivePlatformInfo", string(data)}, " ")
}
