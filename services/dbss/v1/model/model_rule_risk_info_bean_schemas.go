package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleRiskInfoBeanSchemas struct {

	// schema名称
	Schema *string `json:"schema,omitempty"`

	// 表名
	Table *string `json:"table,omitempty"`

	// 列名
	Column *string `json:"column,omitempty"`
}

func (o RuleRiskInfoBeanSchemas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleRiskInfoBeanSchemas struct{}"
	}

	return strings.Join([]string{"RuleRiskInfoBeanSchemas", string(data)}, " ")
}
