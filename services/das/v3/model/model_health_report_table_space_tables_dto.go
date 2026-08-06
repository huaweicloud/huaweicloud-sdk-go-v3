package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HealthReportTableSpaceTablesDto 表信息。
type HealthReportTableSpaceTablesDto struct {

	// 库名。
	DbName *string `json:"db_name,omitempty"`

	// 表名。
	TableName *string `json:"table_name,omitempty"`
}

func (o HealthReportTableSpaceTablesDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportTableSpaceTablesDto struct{}"
	}

	return strings.Join([]string{"HealthReportTableSpaceTablesDto", string(data)}, " ")
}
