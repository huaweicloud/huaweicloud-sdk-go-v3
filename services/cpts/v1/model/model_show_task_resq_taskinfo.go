package model

import (
	"encoding/json"

	"strings"
)

// taskInfo
type ShowTaskResqTaskinfo struct {
	// bench_concurrent

	BenchConcurrent *int32 `json:"bench_concurrent,omitempty"`
	// case_list

	CaseList *[]ShowTaskResqTaskinfoCaseList `json:"case_list,omitempty"`
	// create_time

	CreateTime *string `json:"create_time,omitempty"`
	// description

	Description *string `json:"description,omitempty"`
	// name

	Name *string `json:"name,omitempty"`
	// operate_mode

	OperateMode *int32 `json:"operate_mode,omitempty"`
	// project_id

	ProjectId *int32 `json:"project_id,omitempty"`
	// related_temp_running_data

	RelatedTempRunningData *[]string `json:"related_temp_running_data,omitempty"`
	// run_status

	RunStatus *int32 `json:"run_status,omitempty"`
	// update_time

	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ShowTaskResqTaskinfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTaskResqTaskinfo struct{}"
	}

	return strings.Join([]string{"ShowTaskResqTaskinfo", string(data)}, " ")
}
