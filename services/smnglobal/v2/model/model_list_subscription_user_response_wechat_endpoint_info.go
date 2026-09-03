package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubscriptionUserResponseWechatEndpointInfo struct {

	// 隐去敏感信息的终端地址。
	Endpoint string `json:"endpoint"`
}

func (o ListSubscriptionUserResponseWechatEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionUserResponseWechatEndpointInfo struct{}"
	}

	return strings.Join([]string{"ListSubscriptionUserResponseWechatEndpointInfo", string(data)}, " ")
}
