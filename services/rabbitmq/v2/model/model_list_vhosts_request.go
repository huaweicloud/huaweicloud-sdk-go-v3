package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVhostsRequest Request Object
type ListVhostsRequest struct {

	// **参数解释**： 实例ID。获取方法如下：登录RabbitMQ控制台，在RabbitMQ实例详情页面查找实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 分页查询偏移量，表示从此偏移量开始查询。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页查询单页数量。 **约束限制**： 不涉及。 **取值范围**： 0~50。 **默认取值**： 10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListVhostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVhostsRequest struct{}"
	}

	return strings.Join([]string{"ListVhostsRequest", string(data)}, " ")
}
