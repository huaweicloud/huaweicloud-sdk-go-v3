package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobRequest Request Object
type CreateJobRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *JobInfo `json:"body,omitempty"`
}

func (o CreateJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobRequest struct{}"
	}

	return strings.Join([]string{"CreateJobRequest", string(data)}, " ")
}
