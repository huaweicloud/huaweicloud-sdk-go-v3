package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ForwardingConfigRequestDto **参数说明**：转发配置信息。
type ForwardingConfigRequestDto struct {
	KafkaConfig *KafkaConfigRequestDto `json:"kafka_config,omitempty"`

	MrsKafkaConfig *MrsKafkaConfigRequestDto `json:"mrs_kafka_config,omitempty"`
}

func (o ForwardingConfigRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForwardingConfigRequestDto struct{}"
	}

	return strings.Join([]string{"ForwardingConfigRequestDto", string(data)}, " ")
}
