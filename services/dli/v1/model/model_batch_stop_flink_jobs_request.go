package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopFlinkJobsRequest Request Object
type BatchStopFlinkJobsRequest struct {
	Body *StopFlinkJobsRequestBody `json:"body,omitempty"`
}

func (o BatchStopFlinkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopFlinkJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchStopFlinkJobsRequest", string(data)}, " ")
}
