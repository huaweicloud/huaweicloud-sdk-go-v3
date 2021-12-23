package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopJobRequest struct {
	// 作业名称.

	JobName string `json:"job_name"`
}

func (o StopJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobRequest struct{}"
	}

	return strings.Join([]string{"StopJobRequest", string(data)}, " ")
}
