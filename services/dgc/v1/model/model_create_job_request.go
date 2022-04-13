package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateJobRequest struct {
	Body *JobInfo `json:"body,omitempty"`
}

func (o CreateJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobRequest struct{}"
	}

	return strings.Join([]string{"CreateJobRequest", string(data)}, " ")
}
