package model

import (
	"encoding/json"

	"strings"
)

// UpdateTaskRequestBody
type UpdateTaskRequestBody struct {
	// id

	Id int32 `json:"id"`
	// name

	Name *string `json:"name,omitempty"`
	// description

	Description *string `json:"description,omitempty"`
	// project_id

	ProjectId int32 `json:"project_id"`
	// run_status

	RunStatus *int32 `json:"run_status,omitempty"`
	// run_type

	RunType *int32 `json:"run_type,omitempty"`

	TaskRunInfo *ShowTaskSetResqTaskRunInfo `json:"task_run_info,omitempty"`
	// case_list

	CaseList *[]ShowTaskResqTaskinfoCaseList `json:"case_list,omitempty"`
	// operate_mode

	OperateMode *int32 `json:"operate_mode,omitempty"`
	// bench_concurrent

	BenchConcurrent *int32 `json:"bench_concurrent,omitempty"`
	// related_temp_running_data

	RelatedTempRunningData *[]UpdateTaskRequestBodyRelatedTempRunningData `json:"related_temp_running_data,omitempty"`
}

func (o UpdateTaskRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTaskRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequestBody", string(data)}, " ")
}
