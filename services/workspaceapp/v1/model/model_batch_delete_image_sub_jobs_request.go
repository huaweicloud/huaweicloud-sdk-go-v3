package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteImageSubJobsRequest Request Object
type BatchDeleteImageSubJobsRequest struct {
	Body *BatchDeleteSubJobsReq `json:"body,omitempty"`
}

func (o BatchDeleteImageSubJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteImageSubJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteImageSubJobsRequest", string(data)}, " ")
}
