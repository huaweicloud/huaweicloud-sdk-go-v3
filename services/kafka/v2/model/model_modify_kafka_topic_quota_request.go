package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyKafkaTopicQuotaRequest Request Object
type ModifyKafkaTopicQuotaRequest struct {

	// **参数解释**： 实例ID。获取方法如下：登录Kafka控制台，在Kafka实例详情页面查找实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
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
