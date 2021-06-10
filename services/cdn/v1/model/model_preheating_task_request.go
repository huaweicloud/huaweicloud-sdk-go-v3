package model

import (
	"encoding/json"

	"strings"
)

type PreheatingTaskRequest struct {
	PreheatingTask *PreheatingTaskBody `json:"preheatingTask,omitempty"`
}

func (o PreheatingTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PreheatingTaskRequest struct{}"
	}

	return strings.Join([]string{"PreheatingTaskRequest", string(data)}, " ")
}
