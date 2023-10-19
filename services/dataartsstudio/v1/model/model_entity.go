package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Entity 资产信息
type Entity struct {

	// 数据版本
	Version float32 `json:"version,omitempty"`

	// 关联关系属性，数据类型Map<String, Object>
	RelationshipAttributes *interface{} `json:"relationship_attributes,omitempty"`

	// 父类资产类型
	SuperTypeNames *[]string `json:"super_type_names,omitempty"`

	// 业务属性，数据类型Map<String, Map<String, Object>>
	BusinessAttributes *interface{} `json:"business_attributes,omitempty"`

	// 承担密级和标签的多值对象数据结构，数据结构Map<String, List<Map<String, Object>>>
	MultiAttributes *interface{} `json:"multi_attributes,omitempty"`

	PrivilegeInfo *EntityPrivilegeInfo `json:"privilege_info,omitempty"`

	// 拓展属性，数据结构Map<String, Object>
	ExtendedAttributes *interface{} `json:"extended_attributes,omitempty"`

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

	// 实例化id
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

func (o Entity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Entity struct{}"
	}

	return strings.Join([]string{"Entity", string(data)}, " ")
}
