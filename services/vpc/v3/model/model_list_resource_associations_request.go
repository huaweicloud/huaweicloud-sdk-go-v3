package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceAssociationsRequest Request Object
type ListResourceAssociationsRequest struct {

	// **参数解释**： 资源ID，填写后按照ID进行过滤，支持传入多个ID过滤。 **取值范围**： 弹性网卡/辅助弹性网卡的资源ID。
	ResourceId *[]string `json:"resource_id,omitempty"`

	// **参数解释**： 资源的类型，填写后按照资源类型进行过滤，支持传入多个资源类型过滤。 **取值范围**： - port：弹性网卡。 - subeni：辅助弹性网卡。
	ResourceType *[]string `json:"resource_type,omitempty"`

	// **参数解释**： 资源所属的项目ID，填写后按照项目ID进行过滤。 **取值范围**： 不涉及。
	ResourceProjectId *string `json:"resource_project_id,omitempty"`

	// **参数解释**： 实例ID，填写后按照ID进行过滤，支持传入多个ID过滤。 **取值范围**： ELB实例的ID。
	InstanceId *[]string `json:"instance_id,omitempty"`

	// **参数解释**： 实例的类型，填写后按照实例类型进行过滤，支持传入多个实例类型过滤。 **取值范围**： - ELB：ELB实例。
	InstanceType *[]string `json:"instance_type,omitempty"`

	// **参数解释**： 实例所属的项目ID，填写后按照项目ID进行过滤。 **取值范围**： 不涉及。
	InstanceProjectId *string `json:"instance_project_id,omitempty"`
}

func (o ListResourceAssociationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceAssociationsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceAssociationsRequest", string(data)}, " ")
}
