package model

import (
	"encoding/json"

	"strings"
)

// json
type CreateCaseResqJson struct {
	// task_case_id

	TaskCaseId *int32 `json:"task_case_id,omitempty"`
}

func (o CreateCaseResqJson) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCaseResqJson struct{}"
	}

	return strings.Join([]string{"CreateCaseResqJson", string(data)}, " ")
}
