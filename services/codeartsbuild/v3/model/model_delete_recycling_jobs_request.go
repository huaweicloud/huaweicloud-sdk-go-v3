package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecyclingJobsRequest Request Object
type DeleteRecyclingJobsRequest struct {
	Body *JobsRequestBody `json:"body,omitempty"`
}

func (o DeleteRecyclingJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecyclingJobsRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecyclingJobsRequest", string(data)}, " ")
}
