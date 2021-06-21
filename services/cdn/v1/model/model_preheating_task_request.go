package model

import (
	"encoding/json"

	"strings"
)

type PreheatingTaskRequest struct {
	PreheatingTask *PreheatingTaskRequestPreheatingTask `json:"preheating_task"`
}

func (o PreheatingTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PreheatingTaskRequest struct{}"
	}

	return strings.Join([]string{"PreheatingTaskRequest", string(data)}, " ")
}
