package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketMqConfigsResponse Response Object
type ShowRocketMqConfigsResponse struct {

	// 总数。
	Total float32 `json:"total,omitempty"`

	// 下个分页的offset。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// 上个分页的offset。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

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
