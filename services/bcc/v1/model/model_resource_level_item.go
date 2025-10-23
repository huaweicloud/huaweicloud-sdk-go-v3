package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceLevelItem 资源分级单条数据
type ResourceLevelItem struct {

	// 资源分级ID
	Id *string `json:"id,omitempty"`

	// 资源分级名称
	ResourceLevelName *string `json:"resource_level_name,omitempty"`

	// 资源类型
	Providers *[]string `json:"providers,omitempty"`

	// 资产级别，支持五级 。1公开，2一般，3关键，4重要，5核心
	ResourceLevel *int32 `json:"resource_level,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 账户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 组织ID
	OrganizationId *string `json:"organization_id,omitempty"`

	// 企业项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 应用方式。1：按标签匹配；2：按照资源类型和资源名称正则匹配；3：按照资源ID
	ApplyType *int32 `json:"apply_type,omitempty"`

	// 应用的规则
	ApplyRule *string `json:"apply_rule,omitempty"`

	ComplianceRule *ComplianceRule `json:"compliance_rule,omitempty"`

	// 各类资源数量
	ResourceCountList *[]LevelResourceCount `json:"resource_count_list,omitempty"`
}

func (o ResourceLevelItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceLevelItem struct{}"
	}

	return strings.Join([]string{"ResourceLevelItem", string(data)}, " ")
}
