package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksRequestBody 查询任务列表的请求体
type ListTasksRequestBody struct {

	// 任务类型，包含如下   - cluster_scan：集群扫描   - iac_scan：iac扫描
	TaskType string `json:"task_type"`

	// 本次创建任务的id
	TaskId *string `json:"task_id,omitempty"`

	// 模糊匹配任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 按任务创建时间范围查询时的起始时间
	StartCreateTime *int64 `json:"start_create_time,omitempty"`

	// 按任务创建时间范围查询时的结束时间
	EndCreateTime *int64 `json:"end_create_time,omitempty"`

	// 任务触发类型，包含如下   - manual：手动创建的任务   - schedule：定时任务
	TriggerType *string `json:"trigger_type,omitempty"`

	// 任务状态，包含如下   - ready：等待执行   - running：运行中   - finished：任务结束
	TaskStatus *string `json:"task_status,omitempty"`

	// 排序字段，包含如下   - start_time：任务开始时间
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方式，包含如下   - desc：降序   - asc: 升序
	SortDir *string `json:"sort_dir,omitempty"`

	ClusterScanInfo *ListTasksRequestBodyClusterScanInfo `json:"cluster_scan_info,omitempty"`

	IacScanInfo *ListTasksRequestBodyIacScanInfo `json:"iac_scan_info,omitempty"`
}

func (o ListTasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequestBody struct{}"
	}

	return strings.Join([]string{"ListTasksRequestBody", string(data)}, " ")
}
