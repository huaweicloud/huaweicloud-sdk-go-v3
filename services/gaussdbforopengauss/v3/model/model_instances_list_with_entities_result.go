package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstancesListWithEntitiesResult struct {

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 实例名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**: 资源信息列表。 当参数模板ID对应只读模板时，返回实例下对应的只读节点列表； 其他场景时，返回对应实例信息。
	Entities *[]EntityResult `json:"entities,omitempty"`
}

func (o InstancesListWithEntitiesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesListWithEntitiesResult struct{}"
	}

	return strings.Join([]string{"InstancesListWithEntitiesResult", string(data)}, " ")
}
