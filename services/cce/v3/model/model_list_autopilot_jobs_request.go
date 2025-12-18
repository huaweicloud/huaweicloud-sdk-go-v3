package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutopilotJobsRequest Request Object
type ListAutopilotJobsRequest struct {
}

func (o ListAutopilotJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutopilotJobsRequest struct{}"
	}

	return strings.Join([]string{"ListAutopilotJobsRequest", string(data)}, " ")
}
