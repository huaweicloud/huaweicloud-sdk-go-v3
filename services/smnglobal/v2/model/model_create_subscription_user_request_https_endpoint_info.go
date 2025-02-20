package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubscriptionUserRequestHttpsEndpointInfo struct {

	// 终端地址。必须以“https://”开头。
	Endpoint string `json:"endpoint"`

	// https协议订阅用户的自定义请求头。https协议订阅用户可以自定义请求头。
	Header map[string]string `json:"header,omitempty"`
}

func (o CreateSubscriptionUserRequestHttpsEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequestHttpsEndpointInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequestHttpsEndpointInfo", string(data)}, " ")
}
