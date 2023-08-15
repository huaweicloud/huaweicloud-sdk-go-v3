package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddForwardingConfigsResponse Response Object
type AddForwardingConfigsResponse struct {

	// **参数说明**：转发配置的类型。  **取值范围**：当前仅支持“kafka、mrskafka”。
	ForwardingType *string `json:"forwarding_type,omitempty"`

	KafkaConfig *KafkaConfigResponseDto `json:"kafka_config,omitempty"`

	MrsKafkaConfig *MrsKafkaConfigResponseDto `json:"mrs_kafka_config,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o AddForwardingConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddForwardingConfigsResponse struct{}"
	}

	return strings.Join([]string{"AddForwardingConfigsResponse", string(data)}, " ")
}
