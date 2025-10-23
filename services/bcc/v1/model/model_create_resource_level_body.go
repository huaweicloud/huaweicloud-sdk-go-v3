package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceLevelBody 创建资源分级的body体
type CreateResourceLevelBody struct {

	// 资源分级名称，只能是汉字、大小写字母、下划线、中划线、数字组成，且不能以数字或者中划线或者下划线开头
	ResourceLevelName string `json:"resource_level_name"`

	// 资源类型
	Providers []string `json:"providers"`

	// 资源级别，创建时为5
	ResourceLevel int32 `json:"resource_level"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 应用方式
	ApplyType *int32 `json:"apply_type,omitempty"`

	// 应用的规则, 标签的格式为\"[{\\\"key\\\": \\\"A\\\", \"value\\\":\\\"a\\\"}]\"
	ApplyRule *string `json:"apply_rule,omitempty"`

	// 合规规则的id
	ComplianceRuleId *string `json:"compliance_rule_id,omitempty"`
}

func (o CreateResourceLevelBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceLevelBody struct{}"
	}

	return strings.Join([]string{"CreateResourceLevelBody", string(data)}, " ")
}
