package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKafkaTopicQuotaRequest Request Object
type DeleteKafkaTopicQuotaRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *KafkaTopicQuota `json:"body,omitempty"`
}

func (o DeleteKafkaTopicQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKafkaTopicQuotaRequest struct{}"
	}

	return strings.Join([]string{"DeleteKafkaTopicQuotaRequest", string(data)}, " ")
}
