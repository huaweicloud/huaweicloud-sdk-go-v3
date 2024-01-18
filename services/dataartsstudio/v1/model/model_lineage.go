package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Lineage 关系列表
type Lineage struct {

	// 关系类型。PARENT_CHILD,LOGICAL_PHYSICAL,PK_FK,DATA_FLOW,INSTANCE_OF,JOIN,IS_A,UP_DOWN,ASSOCIATION,WORK_FLOW
	RelationTypeName *string `json:"relation_type_name,omitempty"`

	// 血缘流向，IN,OUT,BOTH
	Direction *string `json:"direction,omitempty"`

	// 节点一资产guid
	Ep1Id *string `json:"ep1_id,omitempty"`

	// 节点一资产类型
	Ep1TypeName *string `json:"ep1_type_name,omitempty"`

	// 节点二资产guid
	Ep2Id *string `json:"ep2_id,omitempty"`

	// 节点二资产类型
	Ep2TypeName *string `json:"ep2_type_name,omitempty"`

	End1 *Entity `json:"end1,omitempty"`

	End2 *Entity `json:"end2,omitempty"`

	// 关系类型。NONE,ONE_TO_TWO,TWO_TO_ONE,BOTH
	PropagateTag *string `json:"propagate_tag,omitempty"`

	// 资产guid
	Guid *string `json:"guid,omitempty"`

	// 资产类型名称
	TypeName *string `json:"type_name,omitempty"`

	// 类型展示名称
	TypeDisplayName *string `json:"type_display_name,omitempty"`

	// 展示名称
	DisplayText *string `json:"display_text,omitempty"`

	// 资产属性，Map<String, Object>
	Attributes *interface{} `json:"attributes,omitempty"`

	// 修改属性列表
	UpdatedAttributes *[]string `json:"updated_attributes,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 主账号id
	DomainId *string `json:"domain_id,omitempty"`

	// 实例id
	InstanceId *[]string `json:"instance_id,omitempty"`

	// 空间id列表
	WorkspaceId *[]string `json:"workspace_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 创建人
	CreatedBy *string `json:"created_by,omitempty"`

	// 修改人
	UpdatedBy *string `json:"updated_by,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o Lineage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Lineage struct{}"
	}

	return strings.Join([]string{"Lineage", string(data)}, " ")
}
