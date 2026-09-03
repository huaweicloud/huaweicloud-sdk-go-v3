package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowFactoryDependInstancesRespDependInstancesInfo struct {

	// 实例ID。
	Id *int64 `json:"id,omitempty"`

	// 作业id。
	JobId *int64 `json:"job_id,omitempty"`

	// 作业名称。
	JobName *string `json:"job_name,omitempty"`

	// 作业路径。
	DirectoryPath *string `json:"directory_path,omitempty"`

	// 实例是否是执行了强制成功。
	ForceSuccess *bool `json:"force_success,omitempty"`

	// 实例是否是执行了忽略失败。
	IgnoreSuccess *bool `json:"ignore_success,omitempty"`

	// 依赖的上游实例ID。
	ParentInstanceIds *[]int64 `json:"parent_instance_ids,omitempty"`

	// 计划开始时间。
	PlanTime *int64 `json:"plan_time,omitempty"`

	// 运行时长，单位：毫秒。 - 当实例是运行中时，运行时长为当前时间减去开始时间； - 当实例运行结束时，运行时长为结束时间减去开始时间；
	RunningTime *int64 `json:"running_time,omitempty"`

	// 开始时间。
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *int64 `json:"end_time,omitempty"`

	// 实例状态。
	Status *string `json:"status,omitempty"`

	// 提交时间。
	SubmitTime *int64 `json:"submit_time,omitempty"`

	// 版本号。
	Version *int32 `json:"version,omitempty"`

	// 所在的工作空间ID。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 所在的工作空间名称。
	WorkspaceName *string `json:"workspace_name,omitempty"`

	// 作业平均执行时长，单位：毫秒。
	AvgExecuteTimeMs *int64 `json:"avg_execute_time_ms,omitempty"`
}

func (o ShowFactoryDependInstancesRespDependInstancesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryDependInstancesRespDependInstancesInfo struct{}"
	}

	return strings.Join([]string{"ShowFactoryDependInstancesRespDependInstancesInfo", string(data)}, " ")
}
