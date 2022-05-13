package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteInstancesRequest struct {
	Body *BatchDeleteInstanceReq `json:"body,omitempty"`
}

func (o BatchDeleteInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstancesRequest", string(data)}, " ")
}
