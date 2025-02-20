package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubscriptionUserRequestSmsEndpointInfo struct {

	// 终端地址。必须是一个电话号码。
	Endpoint string `json:"endpoint"`
}

func (o CreateSubscriptionUserRequestSmsEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequestSmsEndpointInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequestSmsEndpointInfo", string(data)}, " ")
}
