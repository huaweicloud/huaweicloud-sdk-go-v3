package model

import (
	"encoding/json"

	"strings"
)

type CreateCaseResultJson struct {
	// task_case_id

	TaskCaseId *int32 `json:"task_case_id,omitempty"`
}

func (o CreateCaseResultJson) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCaseResultJson struct{}"
	}

	return strings.Join([]string{"CreateCaseResultJson", string(data)}, " ")
}
