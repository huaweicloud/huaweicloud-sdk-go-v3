package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceAssociation ResourceAssociation响应对象
type ResourceAssociation struct {

	// **参数解释**： 资源和实例关联关系的ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 资源ID。 **取值范围**： 不涉及。
	ResourceId string `json:"resource_id"`

	// **参数解释**： 资源类型。 **取值范围**： - port：弹性网卡。 - subeni：辅助弹性网卡。
	ResourceType string `json:"resource_type"`

	// **参数解释**： 资源所属的项目ID。 **取值范围**： 不涉及。
	ResourceProjectId string `json:"resource_project_id"`

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 实例类型。 **取值范围**： - ELB：ELB实例。
	InstanceType string `json:"instance_type"`

	// **参数解释**： 实例所属的项目ID。 **取值范围**： 不涉及。
	InstanceProjectId string `json:"instance_project_id"`
}

func (o ResourceAssociation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceAssociation struct{}"
	}

	return strings.Join([]string{"ResourceAssociation", string(data)}, " ")
}
