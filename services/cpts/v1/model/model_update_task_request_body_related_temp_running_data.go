package model

import (
	"encoding/json"

	"strings"
)

type UpdateTaskRequestBodyRelatedTempRunningData struct {
	// task_run_info_id

	TaskRunInfoId *int32 `json:"task_run_info_id,omitempty"`
	// related_temp_running_id

	RelatedTempRunningId *int32 `json:"related_temp_running_id,omitempty"`
	// temp_id

	TempId *int32 `json:"temp_id,omitempty"`
	// temp_name

	TempName *string `json:"temp_name,omitempty"`
}

func (o UpdateTaskRequestBodyRelatedTempRunningData) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTaskRequestBodyRelatedTempRunningData struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequestBodyRelatedTempRunningData", string(data)}, " ")
}
