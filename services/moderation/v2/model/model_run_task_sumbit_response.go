package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RunTaskSumbitResponse struct {
	Result         *TaskSumbitResponseResult `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o RunTaskSumbitResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunTaskSumbitResponse struct{}"
	}

	return strings.Join([]string{"RunTaskSumbitResponse", string(data)}, " ")
}
