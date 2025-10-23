package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceLevelBase struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 资源等级名称
	ResourceLevelName *string `json:"resource_level_name,omitempty"`

	// 云服务
	Providers *[]string `json:"providers,omitempty"`

	// 重要程度
	ResourceLevel *int32 `json:"resource_level,omitempty"`

	// 账户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 应用方式
	ApplyType *int32 `json:"apply_type,omitempty"`

	// 应用规则
	ApplyRule *string `json:"apply_rule,omitempty"`

	// 合规规则ID
	ComplianceRuleId *string `json:"compliance_rule_id,omitempty"`
}

func (o ResourceLevelBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceLevelBase struct{}"
	}

	return strings.Join([]string{"ResourceLevelBase", string(data)}, " ")
}
