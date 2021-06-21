package model

import (
	"encoding/json"

	"strings"
)

type PreheatingTask struct {
	// 任务id

	Id *string `json:"id,omitempty"`
}

func (o PreheatingTask) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PreheatingTask struct{}"
	}

	return strings.Join([]string{"PreheatingTask", string(data)}, " ")
}
