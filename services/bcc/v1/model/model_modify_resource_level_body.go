package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyResourceLevelBody 修改资源分级的body体
type ModifyResourceLevelBody struct {

	// 资源分级名称，只能是汉字、大小写字母、下划线、中划线、数字组成，且不能以数字或者中划线或者下划线开头
	ResourceLevelName *string `json:"resource_level_name,omitempty"`

	// 资源分级描述。
	Description *string `json:"description,omitempty"`

	// 应用方式
	ApplyType *int32 `json:"apply_type,omitempty"`

	// 应用规则
	ApplyRule *string `json:"apply_rule,omitempty"`

	// 合规规则的id
	ComplianceRuleId *string `json:"compliance_rule_id,omitempty"`

	// 资源等级
	ResourceLevel *int32 `json:"resource_level,omitempty"`
}

func (o ModifyResourceLevelBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyResourceLevelBody struct{}"
	}

	return strings.Join([]string{"ModifyResourceLevelBody", string(data)}, " ")
}
