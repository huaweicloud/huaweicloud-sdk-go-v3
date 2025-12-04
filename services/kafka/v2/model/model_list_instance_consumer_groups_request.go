package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupsRequest Request Object
type ListInstanceConsumerGroupsRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 偏移量。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 不涉及。
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 当次查询返回的最大消费组ID个数。 **约束限制**： 不涉及。 **取值范围**： 1~50。 **默认取值**： 10。
	Limit *string `json:"limit,omitempty"`

	// **参数解释**： 消费组名过滤查询，过滤方式为字段包含过滤。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Group *string `json:"group,omitempty"`
}

func (o ListInstanceConsumerGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupsRequest", string(data)}, " ")
}
