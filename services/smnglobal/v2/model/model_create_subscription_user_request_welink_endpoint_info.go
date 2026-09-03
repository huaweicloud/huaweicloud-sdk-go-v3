package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubscriptionUserRequestWelinkEndpointInfo struct {

	// 终端地址。必须是一个welink的群号。
	Endpoint string `json:"endpoint"`

	// welink协议订阅用户的client_id，从Welink方获取。
	ClientId string `json:"client_id"`

	// welink协议订阅用户的client_secret，从Welink方获取。
	ClientSecret string `json:"client_secret"`
}

func (o CreateSubscriptionUserRequestWelinkEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequestWelinkEndpointInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequestWelinkEndpointInfo", string(data)}, " ")
}
