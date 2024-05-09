package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopFlinkJobsResponse Response Object
type BatchStopFlinkJobsResponse struct {
	Body           *[]FlinkSuccessResponse `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchStopFlinkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopFlinkJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchStopFlinkJobsResponse", string(data)}, " ")
}
