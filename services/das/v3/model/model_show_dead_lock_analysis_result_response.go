package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockAnalysisResultResponse Response Object
type ShowDeadLockAnalysisResultResponse struct {

	// 死锁唯一标识
	DeadLockId *string `json:"dead_lock_id,omitempty"`

	// 分析任务ID
	JobId *string `json:"job_id,omitempty"`

	// 分析任务状态
	Status *string `json:"status,omitempty"`

	// SQL洞察任务ID
	SqlInsightJobId *int64 `json:"sql_insight_job_id,omitempty"`

	// 查询的事务ID
	TransactionId *string `json:"transaction_id,omitempty"`

	// 该事务下SQL记录总数
	Total *int64 `json:"total,omitempty"`

	// SQL详情列表
	SqlList *[]ShowDeadLockAnalysisResultRespSqlList `json:"sql_list,omitempty"`

	// 从死锁事件解析的事务ID列表
	TransactionIds *[]string `json:"transaction_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowDeadLockAnalysisResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockAnalysisResultResponse struct{}"
	}

	return strings.Join([]string{"ShowDeadLockAnalysisResultResponse", string(data)}, " ")
}
