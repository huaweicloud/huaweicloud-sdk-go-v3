package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKafkaTopicQuotaRequest Request Object
type CreateKafkaTopicQuotaRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *KafkaTopicQuota `json:"body,omitempty"`
}

func (o CreateKafkaTopicQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKafkaTopicQuotaRequest struct{}"
	}

	return strings.Join([]string{"CreateKafkaTopicQuotaRequest", string(data)}, " ")
}
