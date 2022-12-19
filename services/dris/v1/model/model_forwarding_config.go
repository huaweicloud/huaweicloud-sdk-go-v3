package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// **参数说明**：转发配置。
type ForwardingConfig struct {

	// **参数说明**：转发配置的类型。  **取值范围**：当前仅支持“kafka、mrskafka”。
	ForwardingType *string `json:"forwarding_type,omitempty"`

	KafkaConfig *KafkaConfigResponseDto `json:"kafka_config,omitempty"`

	MrsKafkaConfig *MrsKafkaConfigResponseDto `json:"mrs_kafka_config,omitempty"`
}

func (o ForwardingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForwardingConfig struct{}"
	}

	return strings.Join([]string{"ForwardingConfig", string(data)}, " ")
}
