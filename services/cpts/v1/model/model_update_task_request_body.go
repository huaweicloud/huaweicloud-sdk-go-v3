package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskRequestBody
type UpdateTaskRequestBody struct {

	// id
	Id int32 `json:"id" xml:"id"`

	// name
	Name string `json:"name" xml:"name"`

	// description
	Description *string `json:"description,omitempty" xml:"description"`

	// project_id
	ProjectId int32 `json:"project_id" xml:"project_id"`

	// run_status
	RunStatus *int32 `json:"run_status,omitempty" xml:"run_status"`

	// run_type
	RunType *int32 `json:"run_type,omitempty" xml:"run_type"`

	TaskRunInfo *TaskRunInfo `json:"task_run_info,omitempty" xml:"task_run_info"`

	// case_list
	CaseList *[]CaseInfo `json:"case_list,omitempty" xml:"case_list"`

	// operate_mode
	OperateMode *int32 `json:"operate_mode,omitempty" xml:"operate_mode"`

	// bench_concurrent
	BenchConcurrent *int32 `json:"bench_concurrent,omitempty" xml:"bench_concurrent"`

	// related_temp_running_data
	RelatedTempRunningData *[]RelatedTempRunningData `json:"related_temp_running_data,omitempty" xml:"related_temp_running_data"`
}

func (o UpdateTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequestBody", string(data)}, " ")
}
