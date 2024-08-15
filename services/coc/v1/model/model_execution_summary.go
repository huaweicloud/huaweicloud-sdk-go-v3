package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecutionSummary 执行相关信息
type ExecutionSummary struct {

	// 工单Id
	OrderId *string `json:"order_id,omitempty"`

	// 脚本执行Id
	JobId *string `json:"job_id,omitempty"`

	// 报告时间
	ReportTime *int64 `json:"report_time,omitempty"`
}

func (o ExecutionSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionSummary struct{}"
	}

	return strings.Join([]string{"ExecutionSummary", string(data)}, " ")
}
