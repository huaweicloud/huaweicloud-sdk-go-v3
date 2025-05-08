package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyKafkaTopicQuotaRequest Request Object
type ModifyKafkaTopicQuotaRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *KafkaTopicQuota `json:"body,omitempty"`
}

func (o ModifyKafkaTopicQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyKafkaTopicQuotaRequest struct{}"
	}

	return strings.Join([]string{"ModifyKafkaTopicQuotaRequest", string(data)}, " ")
}
