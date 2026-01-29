package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserRequest Request Object
type ListUserRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用[查询所有实例列表](ListInstancesDetails.xml)接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 分页查询偏移量。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0。
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 分页查询单页数量。 **约束限制**： 不涉及。 **取值范围**： 0~50。 **默认取值**： 10。
	Limit *string `json:"limit,omitempty"`
}

func (o ListUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserRequest struct{}"
	}

	return strings.Join([]string{"ListUserRequest", string(data)}, " ")
}
