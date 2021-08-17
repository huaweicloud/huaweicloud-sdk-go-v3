package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowCertKeyRequest struct {
	// 迁移任务id

	TaskId string `json:"task_id"`
}

func (o ShowCertKeyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCertKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowCertKeyRequest", string(data)}, " ")
}
