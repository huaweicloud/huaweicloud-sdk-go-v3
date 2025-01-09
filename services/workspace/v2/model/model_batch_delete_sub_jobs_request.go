package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubJobsRequest Request Object
type BatchDeleteSubJobsRequest struct {
	Body *DeleteSubJobsReq `json:"body,omitempty"`
}

func (o BatchDeleteSubJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubJobsRequest", string(data)}, " ")
}
