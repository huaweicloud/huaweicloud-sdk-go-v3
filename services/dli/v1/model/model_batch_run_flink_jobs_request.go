package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRunFlinkJobsRequest Request Object
type BatchRunFlinkJobsRequest struct {
	Body *BatchRunFlinkJobsRequestBody `json:"body,omitempty"`
}

func (o BatchRunFlinkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRunFlinkJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchRunFlinkJobsRequest", string(data)}, " ")
}
