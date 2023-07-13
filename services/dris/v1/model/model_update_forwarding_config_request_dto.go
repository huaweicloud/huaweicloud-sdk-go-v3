package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateForwardingConfigRequestDto 配置更新接口使用的结构，现在支持kafka配置的更新
type UpdateForwardingConfigRequestDto struct {
	KafkaConfig *UpdateKafkaConfigRequestDto `json:"kafka_config,omitempty"`

	MrsKafkaConfig *UpdateMrsKafkaConfigRequestDto `json:"mrs_kafka_config,omitempty"`
}

func (o UpdateForwardingConfigRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateForwardingConfigRequestDto struct{}"
	}

	return strings.Join([]string{"UpdateForwardingConfigRequestDto", string(data)}, " ")
}
