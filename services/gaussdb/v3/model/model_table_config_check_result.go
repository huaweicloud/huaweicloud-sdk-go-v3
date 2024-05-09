package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableConfigCheckResult 表配置校验结果。
type TableConfigCheckResult struct {

	// 表名。
	TableName *string `json:"table_name,omitempty"`

	// 表配置项。
	TableConfig *string `json:"table_config,omitempty"`

	// 校验结果。
	CheckResult *string `json:"check_result,omitempty"`
}

func (o TableConfigCheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableConfigCheckResult struct{}"
	}

	return strings.Join([]string{"TableConfigCheckResult", string(data)}, " ")
}
