package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseDeadLockResponse Response Object
type ParseDeadLockResponse struct {

	// 一键分析死锁日志任务唯一标识符
	JobId *string `json:"job_id,omitempty"`

	// 任务状态
	Status *string `json:"status,omitempty"`

	// SQL洞察任务唯一标识符
	SqlInsightJobId *int64 `json:"sql_insight_job_id,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ParseDeadLockResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseDeadLockResponse struct{}"
	}

	return strings.Join([]string{"ParseDeadLockResponse", string(data)}, " ")
}
