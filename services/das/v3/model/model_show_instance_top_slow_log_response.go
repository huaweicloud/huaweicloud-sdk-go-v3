package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceTopSlowLogResponse Response Object
type ShowInstanceTopSlowLogResponse struct {

	// 采集慢SQL开关
	CollectSlowLog *bool `json:"collect_slow_log,omitempty"`

	// 按执行次数排序的慢SQL列表
	TopExecuteSlowLogs *[]InsTopSlowLogInfo `json:"top_execute_slow_logs,omitempty"`

	// 按平均执行时间排序的慢SQL列表
	TopAvgQueryTimeSlowLogs *[]InsTopSlowLogInfo `json:"top_avg_query_time_slow_logs,omitempty"`

	// 按最大执行时间排序的慢SQL列表
	TopMaxQueryTimeSlowLogs *[]InsTopSlowLogInfo `json:"top_max_query_time_slow_logs,omitempty"`

	// 按扫描返回比排序的慢SQL列表
	RowsExaminedExceeding *[]InsTopSlowLogInfo `json:"rows_examined_exceeding,omitempty"`
	HttpStatusCode        int                  `json:"-"`
}

func (o ShowInstanceTopSlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTopSlowLogResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceTopSlowLogResponse", string(data)}, " ")
}
