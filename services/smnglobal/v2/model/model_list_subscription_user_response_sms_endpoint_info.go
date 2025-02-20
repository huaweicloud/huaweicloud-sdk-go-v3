package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubscriptionUserResponseSmsEndpointInfo struct {

	// 终端地址。
	Endpoint string `json:"endpoint"`
}

func (o ListSubscriptionUserResponseSmsEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionUserResponseSmsEndpointInfo struct{}"
	}

	return strings.Join([]string{"ListSubscriptionUserResponseSmsEndpointInfo", string(data)}, " ")
}
