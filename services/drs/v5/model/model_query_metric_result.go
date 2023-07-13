package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryMetricResult Node节点指标项。
type QueryMetricResult struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 实例ID。
	NodeId *string `json:"node_id,omitempty"`

	// 上报时间。
	TimeStamp *string `json:"time_stamp,omitempty"`

	// CPU使用率。
	CpuUtil *string `json:"cpu_util,omitempty"`

	// 内存使用率。
	MemUtil *string `json:"mem_util,omitempty"`

	// 网络输入吞吐量。
	NetworkIncomingBytesRate *string `json:"network_incoming_bytes_rate,omitempty"`

	// 网络输出吞吐量。
	NetworkOutgoingBytesRate *string `json:"network_outgoing_bytes_rate,omitempty"`

	// 磁盘读吞吐量。
	DiskReadBytesRate *string `json:"disk_read_bytes_rate,omitempty"`

	// 磁盘写吞吐量。
	DiskWriteBytesRate *string `json:"disk_write_bytes_rate,omitempty"`

	// 写目标库频率。
	ApplyRowsRate *string `json:"apply_rows_rate,omitempty"`

	// DML TPS。
	ApplyTransactionsRate *string `json:"apply_transactions_rate,omitempty"`

	// DDL TPS。
	ApplyDdlRate *string `json:"apply_ddl_rate,omitempty"`

	// 事务平均执行时间。
	ApplyAverageExecuteTime *string `json:"apply_average_execute_time,omitempty"`

	// 事务平均提交时间。
	ApplyAverageCommitTime *string `json:"apply_average_commit_time,omitempty"`

	// 同步状态。
	ApplyCurrentState *string `json:"apply_current_state,omitempty"`
}

func (o QueryMetricResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryMetricResult struct{}"
	}

	return strings.Join([]string{"QueryMetricResult", string(data)}, " ")
}
