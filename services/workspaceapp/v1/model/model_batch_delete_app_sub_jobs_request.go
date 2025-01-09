package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppSubJobsRequest Request Object
type BatchDeleteAppSubJobsRequest struct {
	Body *BatchDeleteSubJobsReq `json:"body,omitempty"`
}

func (o BatchDeleteAppSubJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppSubJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppSubJobsRequest", string(data)}, " ")
}
