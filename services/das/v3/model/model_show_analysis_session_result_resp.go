package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAnalysisSessionResultResp struct {

	// 分析的会话总数
	TotalCount int64 `json:"total_count"`

	// 状态持续时间长TOP会话列表
	TopStateDuration []ShowAnalysisSessionResultRespTopStateDuration `json:"top_state_duration"`

	// 事务持续时间长TOP会话列表
	TopTransactionDuration []ShowAnalysisSessionResultRespTopTransactionDuration `json:"top_transaction_duration"`

	// SQL模板列表
	SqlTemplates []ShowAnalysisSessionResultRespSqlTemplates `json:"sql_templates"`
}

func (o ShowAnalysisSessionResultResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAnalysisSessionResultResp struct{}"
	}

	return strings.Join([]string{"ShowAnalysisSessionResultResp", string(data)}, " ")
}
