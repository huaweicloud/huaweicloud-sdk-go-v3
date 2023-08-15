package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchSslRequest struct {

	// SSL数据加密开关设置。  - true: 开启SSL数据加密。 - false: 关闭SSL数据加密。
	SslOption bool `json:"ssl_option"`
}

func (o SwitchSslRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSslRequest struct{}"
	}

	return strings.Join([]string{"SwitchSslRequest", string(data)}, " ")
}
