package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreRecyclingJobsRequest Request Object
type RestoreRecyclingJobsRequest struct {
	Body *JobsRequestBody `json:"body,omitempty"`
}

func (o RestoreRecyclingJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreRecyclingJobsRequest struct{}"
	}

	return strings.Join([]string{"RestoreRecyclingJobsRequest", string(data)}, " ")
}
