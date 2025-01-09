package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteJobsRequest Request Object
type BatchDeleteJobsRequest struct {
	Body *BatchOperateJobsReq `json:"body,omitempty"`
}

func (o BatchDeleteJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobsRequest", string(data)}, " ")
}
