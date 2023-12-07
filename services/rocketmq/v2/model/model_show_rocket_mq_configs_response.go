package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketMqConfigsResponse Response Object
type ShowRocketMqConfigsResponse struct {

	// RocketMQ配置。
	RocketmqConfigs *[]RocketMqConfigResp `json:"rocketmq_configs,omitempty"`
	HttpStatusCode  int                   `json:"-"`
}

func (o ShowRocketMqConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketMqConfigsResponse struct{}"
	}

	return strings.Join([]string{"ShowRocketMqConfigsResponse", string(data)}, " ")
}
