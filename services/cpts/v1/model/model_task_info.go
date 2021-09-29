package model

import (
	"encoding/json"

	"strings"
)

type TaskInfo struct {
	// bench_concurrent

	BenchConcurrent *int32 `json:"bench_concurrent,omitempty"`
	// case_list

	CaseList *[]CaseInfo `json:"case_list,omitempty"`
}

func (o TaskInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TaskInfo struct{}"
	}

	return strings.Join([]string{"TaskInfo", string(data)}, " ")
}
