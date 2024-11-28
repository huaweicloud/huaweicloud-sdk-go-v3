package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLivePlatformReq 自定义直播平台创建请求
type CreateLivePlatformReq struct {

	// 直播平台名称
	Name string `json:"name"`

	AuthConfig *CustomPlatformAuthConfig `json:"auth_config"`

	// 自定义直播平台回调配置。同一种类型仅保留一个配置，如果配置多个会随机保存一个。
	CallbackConfig []StandardPlatformApiConfig `json:"callback_config"`
}

func (o CreateLivePlatformReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLivePlatformReq struct{}"
	}

	return strings.Join([]string{"CreateLivePlatformReq", string(data)}, " ")
}
