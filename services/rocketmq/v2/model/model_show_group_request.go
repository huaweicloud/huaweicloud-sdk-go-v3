package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupRequest Request Object
type ShowGroupRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Group string `json:"group"`
}

func (o ShowGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupRequest", string(data)}, " ")
}
