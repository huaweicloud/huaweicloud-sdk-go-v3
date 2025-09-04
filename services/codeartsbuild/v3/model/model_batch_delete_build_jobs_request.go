package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBuildJobsRequest Request Object
type BatchDeleteBuildJobsRequest struct {
	Body *BatchDeleteBuildJobsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteBuildJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBuildJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteBuildJobsRequest", string(data)}, " ")
}
