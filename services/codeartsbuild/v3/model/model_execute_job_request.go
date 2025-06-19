package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteJobRequest Request Object
type ExecuteJobRequest struct {
	Body *RunJobRequestBody `json:"body,omitempty"`
}

func (o ExecuteJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteJobRequest struct{}"
	}

	return strings.Join([]string{"ExecuteJobRequest", string(data)}, " ")
}
