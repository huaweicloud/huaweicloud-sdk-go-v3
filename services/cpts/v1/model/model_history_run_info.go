package model

import (
	"encoding/json"

	"strings"
)

type HistoryRunInfo struct {
	// name

	Name *string `json:"name,omitempty"`
	// run_id

	RunId *int32 `json:"run_id,omitempty"`
	// run_type

	RunType *int32 `json:"run_type,omitempty"`
	// start_time

	StartTime *string `json:"start_time,omitempty"`
	// continue_time

	ContinueTime *int32 `json:"continue_time,omitempty"`
	// temp_names

	TempNames *[]TempName `json:"temp_names,omitempty"`
}

func (o HistoryRunInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HistoryRunInfo struct{}"
	}

	return strings.Join([]string{"HistoryRunInfo", string(data)}, " ")
}
