package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubscriptionUserRequestWechatEndpointInfo struct {

	// 终端地址。必须是一个微信群机器人的地址。
	Endpoint string `json:"endpoint"`
}

func (o CreateSubscriptionUserRequestWechatEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequestWechatEndpointInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequestWechatEndpointInfo", string(data)}, " ")
}
