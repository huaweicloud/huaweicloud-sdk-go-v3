package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceNodesRequest Request Object
type ShowInstanceNodesRequest struct {

	// **参数解释**： 消息引擎类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Engine string `json:"engine"`

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 查询数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，表示从此偏移量开始查询，offset大于等于0。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowInstanceNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceNodesRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceNodesRequest", string(data)}, " ")
}
