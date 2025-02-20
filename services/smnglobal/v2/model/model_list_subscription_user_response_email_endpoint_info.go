package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubscriptionUserResponseEmailEndpointInfo struct {

	// 终端地址。
	Endpoint string `json:"endpoint"`
}

func (o ListSubscriptionUserResponseEmailEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionUserResponseEmailEndpointInfo struct{}"
	}

	return strings.Join([]string{"ListSubscriptionUserResponseEmailEndpointInfo", string(data)}, " ")
}
