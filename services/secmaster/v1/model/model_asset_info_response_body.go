package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetInfoResponseBody 操作连接信息参数对象
type AssetInfoResponseBody struct {

	// 操作连接ID
	Id *string `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 操作连接名称
	Name *string `json:"name,omitempty"`

	// 操作连接所属的插件id
	ComponentId *string `json:"component_id,omitempty"`

	// 操作连接所属的插件id名称
	ComponentName *string `json:"component_name,omitempty"`

	// 插件版本ID
	ComponentVersionId *string `json:"component_version_id,omitempty"`

	// 操作连接类型
	Type *string `json:"type,omitempty"`

	// 操作连接状态（SUCCESS/FAILED）
	Status *string `json:"status,omitempty"`

	// 具体的操作连接配置字符串，根据插件的操作连接schema配置对应字段值
	Config *string `json:"config,omitempty"`

	// 操作连接描述
	Description *string `json:"description,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 更新者ID
	ModifierId *string `json:"modifier_id,omitempty"`

	// 更新者名称
	ModifierName *string `json:"modifier_name,omitempty"`

	// 下发应急策略时的所属的防线分类
	DefenseType *string `json:"defense_type,omitempty"`

	// 下发应急策略时的IAM项目ID
	TargetProjectId *string `json:"target_project_id,omitempty"`

	// 下发应急策略时的IAM项目名称
	TargetProjectName *string `json:"target_project_name,omitempty"`

	// 下发应急策略时的企业项目ID
	TargetEnterpriseId *string `json:"target_enterprise_id,omitempty"`

	// 下发应急策略时的企业项目名称
	TargetEnterpriseName *string `json:"target_enterprise_name,omitempty"`

	// 下发应急策略时的区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 下发应急策略时的区域名称
	RegionName *string `json:"region_name,omitempty"`

	// 是否可删除，操作连接有流程在使用时，返回false
	CanBeDeleted *bool `json:"can_be_deleted,omitempty"`
}

func (o AssetInfoResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetInfoResponseBody struct{}"
	}

	return strings.Join([]string{"AssetInfoResponseBody", string(data)}, " ")
}
