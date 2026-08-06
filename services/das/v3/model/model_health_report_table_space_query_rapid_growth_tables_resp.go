package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportTableSpaceQueryRapidGrowthTablesResp struct {

	// 表列表。
	Tables *[]HealthReportTableSpaceTopDataDto `json:"tables,omitempty"`

	// 阈值。
	Threshold *int64 `json:"threshold,omitempty"`

	// 上次诊断时间。
	LastDiagnoseTimestamp *int64 `json:"last_diagnose_timestamp,omitempty"`
}

func (o HealthReportTableSpaceQueryRapidGrowthTablesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportTableSpaceQueryRapidGrowthTablesResp struct{}"
	}

	return strings.Join([]string{"HealthReportTableSpaceQueryRapidGrowthTablesResp", string(data)}, " ")
}
