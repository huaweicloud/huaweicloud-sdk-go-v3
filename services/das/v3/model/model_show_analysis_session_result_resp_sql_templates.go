package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAnalysisSessionResultRespSqlTemplates struct {

	// SQL模板
	SqlTemplate string `json:"sql_template"`

	// 数据库名
	DatabaseName string `json:"database_name"`

	// 总执行次数
	TotalCount int64 `json:"total_count"`

	// 当前模板下状态持续时间长TOP会话列表
	TopStateDurationList []ShowAnalysisSessionResultRespTopStateDuration `json:"top_state_duration_list"`

	// 当前模板下事务持续时间长TOP会话列表
	TopTransactionDurationList []ShowAnalysisSessionResultRespTopStateDuration `json:"top_transaction_duration_list"`
}

func (o ShowAnalysisSessionResultRespSqlTemplates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAnalysisSessionResultRespSqlTemplates struct{}"
	}

	return strings.Join([]string{"ShowAnalysisSessionResultRespSqlTemplates", string(data)}, " ")
}
