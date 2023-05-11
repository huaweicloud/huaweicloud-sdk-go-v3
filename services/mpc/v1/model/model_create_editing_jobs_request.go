package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEditingJobsRequest struct {
	Body *CreateEditingJobsReq `json:"body,omitempty"`
}

func (o CreateEditingJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditingJobsRequest struct{}"
	}

	return strings.Join([]string{"CreateEditingJobsRequest", string(data)}, " ")
}
