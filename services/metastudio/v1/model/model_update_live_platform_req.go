package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLivePlatformReq 自定义直播平台创建请求
type UpdateLivePlatformReq struct {

	// 直播平台名称
	Name string `json:"name"`

	AuthConfig *UpdateCustomPlatformAuthConfig `json:"auth_config,omitempty"`

	// 自定义直播平台回调配置。同一种类型仅保留一个配置，如果配置多个会随机保存一个。
	CallbackConfig []StandardPlatformApiConfig `json:"callback_config"`
}

func (o UpdateLivePlatformReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLivePlatformReq struct{}"
	}

	return strings.Join([]string{"UpdateLivePlatformReq", string(data)}, " ")
}
