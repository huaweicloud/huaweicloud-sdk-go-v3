package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopicProducersRequest Request Object
type ListTopicProducersRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 主题。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topic string `json:"topic"`

	// **参数解释**： 偏移量，表示查询该偏移量后面的记录。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 查询返回记录的数量限制。 **约束限制**： 不涉及。 **取值范围**： 1~50。 **默认取值**： 不涉及。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTopicProducersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicProducersRequest struct{}"
	}

	return strings.Join([]string{"ListTopicProducersRequest", string(data)}, " ")
}
