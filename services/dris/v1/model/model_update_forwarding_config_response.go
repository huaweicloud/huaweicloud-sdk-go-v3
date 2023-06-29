package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateForwardingConfigResponse Response Object
type UpdateForwardingConfigResponse struct {

	// **参数说明**：转发配置的类型。  **取值范围**：当前仅支持“kafka、mrskafka”。
	ForwardingType *string `json:"forwarding_type,omitempty"`

	KafkaConfig *KafkaConfigResponseDto `json:"kafka_config,omitempty"`

	MrsKafkaConfig *MrsKafkaConfigResponseDto `json:"mrs_kafka_config,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o UpdateForwardingConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateForwardingConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateForwardingConfigResponse", string(data)}, " ")
}
