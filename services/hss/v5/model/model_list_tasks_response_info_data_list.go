package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksResponseInfoDataList 任务信息
type ListTasksResponseInfoDataList struct {

	// 本次创建任务的id
	TaskId *string `json:"task_id,omitempty"`

	// 任务类型，包含如下   - cluster_scan：集群扫描   - iac_scan：iac扫描
	TaskType *string `json:"task_type,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 任务触发类型，包含如下   - manual：手动创建的任务   - schedule：定时任务
	TriggerType *string `json:"trigger_type,omitempty"`

	// 任务状态，包含如下   - ready：等待执行   - running：运行中   - finished：任务结束
	TaskStatus *string `json:"task_status,omitempty"`

	// 任务开始的时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 任务结束的时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 预计当前任务结束还需要的时间，单位为分钟
	EstimatedTime *int32 `json:"estimated_time,omitempty"`

	ClusterScanInfo *ListTasksResponseInfoClusterScanInfo `json:"cluster_scan_info,omitempty"`

	IacScanInfo *ListTasksResponseInfoIacScanInfo `json:"iac_scan_info,omitempty"`
}

func (o ListTasksResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"ListTasksResponseInfoDataList", string(data)}, " ")
}
