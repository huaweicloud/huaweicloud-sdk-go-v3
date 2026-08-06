package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HealthReportTableSpaceQuerySpecialTablesResp 特殊表列表。
type HealthReportTableSpaceQuerySpecialTablesResp struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 最近诊断时间。
	LastDiagnoseTimestamp *int64 `json:"last_diagnose_timestamp,omitempty"`

	// 库表信息列表。
	Tables *[]HealthReportTableSpaceTablesDto `json:"tables,omitempty"`
}

func (o HealthReportTableSpaceQuerySpecialTablesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportTableSpaceQuerySpecialTablesResp struct{}"
	}

	return strings.Join([]string{"HealthReportTableSpaceQuerySpecialTablesResp", string(data)}, " ")
}
