package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaTopicQuotaRequest Request Object
type ShowKafkaTopicQuotaRequest struct {

	// **参数解释**： 实例ID。获取方法如下：登录Kafka控制台，在Kafka实例详情页面查找实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 查询类型。 **约束限制**： 查询类型必须以字母开头且只支持大小写字母、中横线、下划线以及数字。 **取值范围**： 默认为topic **默认取值**： topic
	Type *string `json:"type,omitempty"`

	// **参数解释**： 每一页显示的流控数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Limit *string `json:"limit,omitempty"`

	// **参数解释**： 页数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 查询关键字。 **约束限制**： 查询关键字必须以字母开头且只支持大小写字母、中横线、下划线以及数字。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Keyword *string `json:"keyword,omitempty"`
}

func (o ShowKafkaTopicQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaTopicQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaTopicQuotaRequest", string(data)}, " ")
}
