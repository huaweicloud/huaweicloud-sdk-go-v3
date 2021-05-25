package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobInstanceRequest struct {
	// 作业名称.

	JobName string `json:"job_name"`
	// 作业实例id.

	InstanceId string `json:"instance_id"`
}

func (o ShowJobInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowJobInstanceRequest", string(data)}, " ")
}
