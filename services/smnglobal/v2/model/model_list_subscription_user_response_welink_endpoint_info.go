package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubscriptionUserResponseWelinkEndpointInfo struct {

	// 终端地址。
	Endpoint string `json:"endpoint"`

	// welink协议订阅用户的client_id，从Welink方获取。
	ClientId string `json:"client_id"`
}

func (o ListSubscriptionUserResponseWelinkEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionUserResponseWelinkEndpointInfo struct{}"
	}

	return strings.Join([]string{"ListSubscriptionUserResponseWelinkEndpointInfo", string(data)}, " ")
}
