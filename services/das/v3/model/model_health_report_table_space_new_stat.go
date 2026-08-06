package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportTableSpaceNewStat struct {

	// top库列表。
	DatabaseTopResp *[]HealthReportTableSpaceQuerySpaceTopResp `json:"database_top_resp,omitempty"`

	// top表列表。
	TableTopResp *[]HealthReportTableSpaceQuerySpaceTopResp `json:"table_top_resp,omitempty"`

	// 异常增长表列表。
	RapidGrowthTablesResp *[]HealthReportTableSpaceQueryRapidGrowthTablesResp `json:"rapid_growth_tables_resp,omitempty"`

	// 无主键表列表。
	NoPrimaryTablesResp *[]HealthReportTableSpaceQuerySpecialTablesResp `json:"no_primary_tables_resp,omitempty"`

	// 无索引表列表。
	NoIndexTablesResp *[]HealthReportTableSpaceQuerySpecialTablesResp `json:"no_index_tables_resp,omitempty"`
}

func (o HealthReportTableSpaceNewStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportTableSpaceNewStat struct{}"
	}

	return strings.Join([]string{"HealthReportTableSpaceNewStat", string(data)}, " ")
}
