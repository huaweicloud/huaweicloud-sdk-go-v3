package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubscriptionUserRequestEmailEndpointInfo struct {

	// 终端地址。必须是邮件地址。
	Endpoint string `json:"endpoint"`
}

func (o CreateSubscriptionUserRequestEmailEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequestEmailEndpointInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequestEmailEndpointInfo", string(data)}, " ")
}
