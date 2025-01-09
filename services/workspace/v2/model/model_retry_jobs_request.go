package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryJobsRequest Request Object
type RetryJobsRequest struct {
	Body *BatchOperateJobsReq `json:"body,omitempty"`
}

func (o RetryJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryJobsRequest struct{}"
	}

	return strings.Join([]string{"RetryJobsRequest", string(data)}, " ")
}
