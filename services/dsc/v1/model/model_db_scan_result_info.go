package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DbScanResultInfo struct {
	// 任务ID

	TaskId *string `json:"task_id,omitempty"`
	// 数据库名称

	DbName *string `json:"db_name,omitempty"`
	// 表ID

	TableId *string `json:"table_id,omitempty"`
	// 表名称

	TableName *string `json:"table_name,omitempty"`
	// 风险等级

	RiskLevel *int32 `json:"risk_level,omitempty"`
	// 匹配到的规则

	SensitiveDataType *[]string `json:"sensitive_data_type,omitempty"`
	// 表中各列匹配到的规则

	MatchInfo *[]DbMatchInfo `json:"match_info,omitempty"`
}

func (o DbScanResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbScanResultInfo struct{}"
	}

	return strings.Join([]string{"DbScanResultInfo", string(data)}, " ")
}
