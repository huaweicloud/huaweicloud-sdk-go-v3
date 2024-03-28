package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubJob 子作业JobData信息，对应spark的JobData。
type SubJob struct {

	// 子作业ID，对应开源spark JobData的jobId。
	Id *int32 `json:"id,omitempty"`

	// 子作业name，对应开源spark JobData的name。
	Name *string `json:"name,omitempty"`

	// 子作业description，对应开源spark JobData的description。
	Description *string `json:"description,omitempty"`

	// 子作业submission_time，对应开源spark JobData的submissionTime。
	SubmissionTime *string `json:"submission_time,omitempty"`

	// 子作业completion_time，对应开源spark JobData的completionTime。
	CompletionTime *string `json:"completion_time,omitempty"`

	// 子作业stage_ids，对应开源spark JobData的stageIds。
	StageIds *[]int32 `json:"stage_ids,omitempty"`

	// 对应DLI的作业ID，对应开源spark JobData的jobGroup。
	JobGroup *string `json:"job_group,omitempty"`

	// 子作业状态，对应开源spark JobData的status。
	Status *string `json:"status,omitempty"`

	// 子作业task的个数，对应开源spark JobData的numTasks。
	NumTasks *int32 `json:"num_tasks,omitempty"`

	// 子作业正在运行的task个数，对应开源spark JobData的numActiveTasks。
	NumActiveTasks *int32 `json:"num_active_tasks,omitempty"`

	// 子作业已经完成的task个数，对应开源spark JobData的numCompletedTasks。
	NumCompletedTasks *int32 `json:"num_completed_tasks,omitempty"`

	// 子作业跳过的task个数，对应开源spark JobData的numSkippedTasks。
	NumSkippedTasks *int32 `json:"num_skipped_tasks,omitempty"`

	// 子作业跳失败的task个数，对应开源spark JobData的numFailedTasks。
	NumFailedTasks *int32 `json:"num_failed_tasks,omitempty"`

	// 子作业kill掉的task个数，对应开源spark JobData的numKilledTasks。
	NumKilledTasks *int32 `json:"num_killed_tasks,omitempty"`

	// 子作业完成指数，对应开源spark JobData的numCompletedIndices。
	NumCompletedIndices *int32 `json:"num_completed_indices,omitempty"`

	// 子作业正在运行的stage个数，对应开源spark JobData的numActiveStages。
	NumActiveStages *int32 `json:"num_active_stages,omitempty"`

	// 子作业已经完成的stage个数，对应开源spark JobData的numCompletedStages。
	NumCompletedStages *int32 `json:"num_completed_stages,omitempty"`

	// 子作业跳过的stage个数，对应开源spark JobData的numSkippedStages。
	NumSkippedStages *int32 `json:"num_skipped_stages,omitempty"`

	// 子作业失败的stage个数，对应开源spark JobData的numFailedStages。
	NumFailedStages *int32 `json:"num_failed_stages,omitempty"`

	// 子作业killed_tasks_summary，对应开源spark JobData的killedTasksSummary。
	KilledTasksSummary map[string]int32 `json:"killed_tasks_summary,omitempty"`
}

func (o SubJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJob struct{}"
	}

	return strings.Join([]string{"SubJob", string(data)}, " ")
}
