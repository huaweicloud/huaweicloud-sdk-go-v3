package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchSslRequestBody struct {

	// ss开关选项。 -“on”，表示NoSQL实例默认开启SSL连接。 -“off”，表示NoSQL实例默认不启用SSL连接。
	SslOption string `json:"ssl_option"`
}

func (o SwitchSslRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSslRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchSslRequestBody", string(data)}, " ")
}
