package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobMonitorInfo 作业监控信息。
type JobMonitorInfo struct {

	// 任务更新时间，毫秒时间戳。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 作业源信息。
	NodeId *string `json:"node_id,omitempty"`

	// 作业消费位点。
	ConsumePosition *string `json:"consume_position,omitempty"`

	// 起始位点。
	OriginPosition *string `json:"origin_position,omitempty"`

	// 作业全量增量运行状态。 - INITIALIZING：初始化 - BINLOG：增量同步 - SNAPSHOT：全量同步
	RunningStatus *string `json:"running_status,omitempty"`

	// 单节点聚合后的监控指标。
	TotalTaskProps *interface{} `json:"total_task_props,omitempty"`

	// 连接列表。
	TaskInfo *[]MonitorTaskInfo `json:"task_info,omitempty"`

	SnapshotProgress *SnapshotProgressInfo `json:"snapshot_progress,omitempty"`
}

func (o JobMonitorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobMonitorInfo struct{}"
	}

	return strings.Join([]string{"JobMonitorInfo", string(data)}, " ")
}
