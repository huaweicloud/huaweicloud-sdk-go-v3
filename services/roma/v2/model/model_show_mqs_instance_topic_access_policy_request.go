package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMqsInstanceTopicAccessPolicyRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// topic名称。

	TopicName string `json:"topic_name"`
	// 分页查询偏移量。

	Offset *string `json:"offset,omitempty"`
	// 分页查询大小。

	Limit *string `json:"limit,omitempty"`
}

func (o ShowMqsInstanceTopicAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMqsInstanceTopicAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowMqsInstanceTopicAccessPolicyRequest", string(data)}, " ")
}
