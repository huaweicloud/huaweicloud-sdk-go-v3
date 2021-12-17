package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopJobRequest struct {
	// 作业ID

	JobId string `json:"job_id"`
	// 停止作业触发savepoint

	TriggerSavepoint *bool `json:"trigger_savepoint,omitempty"`
}

func (o StopJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobRequest struct{}"
	}

	return strings.Join([]string{"StopJobRequest", string(data)}, " ")
}
