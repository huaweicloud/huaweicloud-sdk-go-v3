package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaseOpsKeyViewRequestBody struct {

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 节点ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**： 组件ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ComponentId string `json:"component_id"`
}

func (o BaseOpsKeyViewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseOpsKeyViewRequestBody struct{}"
	}

	return strings.Join([]string{"BaseOpsKeyViewRequestBody", string(data)}, " ")
}
