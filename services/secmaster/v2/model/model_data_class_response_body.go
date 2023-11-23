package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataClassResponseBody 数据类详情
type DataClassResponseBody struct {

	// 数据类ID
	Id *string `json:"id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 修改者ID
	ModifierId *string `json:"modifier_id,omitempty"`

	// 修改这名称
	ModifierName *string `json:"modifier_name,omitempty"`

	// 订阅包版本
	CloudPackVersion *string `json:"cloud_pack_version,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// domain id
	DomainId *string `json:"domain_id,omitempty"`

	// 数据类名称
	Name *string `json:"name,omitempty"`

	// 数据类业务编码
	BusinessCode *string `json:"business_code,omitempty"`

	// 数据类描述
	Description *string `json:"description,omitempty"`

	// 是否内置，true内置，false非内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 父级id
	ParentId *string `json:"parent_id,omitempty"`

	// 子类型数量
	TypeNum float32 `json:"type_num,omitempty"`
}

func (o DataClassResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassResponseBody struct{}"
	}

	return strings.Join([]string{"DataClassResponseBody", string(data)}, " ")
}
