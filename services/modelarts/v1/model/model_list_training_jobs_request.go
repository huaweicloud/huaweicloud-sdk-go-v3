package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrainingJobsRequest Request Object
type ListTrainingJobsRequest struct {
	Body *JobSearches `json:"body,omitempty"`
}

func (o ListTrainingJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingJobsRequest struct{}"
	}

	return strings.Join([]string{"ListTrainingJobsRequest", string(data)}, " ")
}
