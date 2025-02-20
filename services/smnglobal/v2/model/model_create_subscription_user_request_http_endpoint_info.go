package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubscriptionUserRequestHttpEndpointInfo struct {

	// 终端地址。必须以“http://”开头。
	Endpoint string `json:"endpoint"`

	// http协议订阅用户的自定义请求头。http协议订阅用户可以自定义请求头。
	Header map[string]string `json:"header,omitempty"`
}

func (o CreateSubscriptionUserRequestHttpEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequestHttpEndpointInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequestHttpEndpointInfo", string(data)}, " ")
}
