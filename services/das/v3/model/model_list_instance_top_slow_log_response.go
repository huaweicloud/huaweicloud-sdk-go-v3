package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceTopSlowLogResponse Response Object
type ListInstanceTopSlowLogResponse struct {

	// 执行次数列表
	TopExecuteSlowLogs *[]TopInstanceSlowLogTopExecuteSlowLogs `json:"top_execute_slow_logs,omitempty"`

	// 平均执行时间列表
	TopAvgQueryTimeSlowLogs *[]TopInstanceSlowLogTopAvgQueryTimeSlowLogs `json:"top_avg_query_time_slow_logs,omitempty"`

	// 最大执行时间列表
	TopMaxQueryTimeSlowLogs *[]TopInstanceSlowLogTopExecuteSlowLogs `json:"top_max_query_time_slow_logs,omitempty"`

	// 扫描返回比列表
	RowsExaminedExceeding *[]TopInstanceSlowLogRowsExaminedExceeding `json:"rows_examined_exceeding,omitempty"`
	HttpStatusCode        int                                        `json:"-"`
}

func (o ListInstanceTopSlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTopSlowLogResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceTopSlowLogResponse", string(data)}, " ")
}
