package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopFlinkJobsRequest Request Object
type StopFlinkJobsRequest struct {
	Body *StopFlinkJobsRequestBody `json:"body,omitempty"`
}

func (o StopFlinkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopFlinkJobsRequest struct{}"
	}

	return strings.Join([]string{"StopFlinkJobsRequest", string(data)}, " ")
}
