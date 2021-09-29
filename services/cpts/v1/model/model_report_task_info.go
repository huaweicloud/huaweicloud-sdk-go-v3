package model

import (
	"encoding/json"

	"strings"
)

type ReportTaskInfo struct {
	// vum

	Vum *int32 `json:"vum,omitempty"`
}

func (o ReportTaskInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReportTaskInfo struct{}"
	}

	return strings.Join([]string{"ReportTaskInfo", string(data)}, " ")
}
