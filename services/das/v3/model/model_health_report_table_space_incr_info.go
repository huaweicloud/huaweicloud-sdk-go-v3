package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportTableSpaceIncrInfo struct {

	// 数据库名。
	Database string `json:"database"`

	// 表名。
	Table string `json:"table"`

	// 增长量。
	Increment int64 `json:"increment"`

	// 统计分析是否成功。
	AnalyzeSuccess bool `json:"analyze_success"`

	// 错误信息。
	ErrorMessage string `json:"error_message"`
}

func (o HealthReportTableSpaceIncrInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportTableSpaceIncrInfo struct{}"
	}

	return strings.Join([]string{"HealthReportTableSpaceIncrInfo", string(data)}, " ")
}
