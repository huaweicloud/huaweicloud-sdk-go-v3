package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchProxySslRequest 开关数据库代理SSL加密请求体。
type SwitchProxySslRequest struct {

	// SSL数据加密开关设置。    取值范围： - true: 开启SSL数据加密。 - false: 关闭SSL数据加密。
	SslOption bool `json:"ssl_option"`
}

func (o SwitchProxySslRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchProxySslRequest struct{}"
	}

	return strings.Join([]string{"SwitchProxySslRequest", string(data)}, " ")
}
