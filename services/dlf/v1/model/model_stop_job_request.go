package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopJobRequest Request Object
type StopJobRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

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
