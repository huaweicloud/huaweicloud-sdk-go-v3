package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceTopicsRequest Request Object
type ListInstanceTopicsRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 偏移量，表示从此偏移量开始查询。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 不涉及。
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 当次查询返回的最大实例个数。 **约束限制**： 不涉及。 **取值范围**： 大于等于0，小于等于200。 **默认取值**： 50。
	Limit *string `json:"limit,omitempty"`
}

func (o ListInstanceTopicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTopicsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceTopicsRequest", string(data)}, " ")
}
