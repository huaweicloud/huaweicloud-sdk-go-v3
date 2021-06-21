package model

import (
	"encoding/json"

	"strings"
)

type RefreshTask struct {
	// 任务id

	Id *string `json:"id,omitempty"`
}

func (o RefreshTask) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RefreshTask struct{}"
	}

	return strings.Join([]string{"RefreshTask", string(data)}, " ")
}
