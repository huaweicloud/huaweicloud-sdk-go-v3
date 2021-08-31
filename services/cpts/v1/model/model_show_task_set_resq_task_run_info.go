package model

import (
	"encoding/json"

	"strings"
)

// task_run_info
type ShowTaskSetResqTaskRunInfo struct {
	// id

	Id *int32 `json:"id,omitempty"`
	// run_type

	RunType *int32 `json:"run_type,omitempty"`
}

func (o ShowTaskSetResqTaskRunInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTaskSetResqTaskRunInfo struct{}"
	}

	return strings.Join([]string{"ShowTaskSetResqTaskRunInfo", string(data)}, " ")
}
