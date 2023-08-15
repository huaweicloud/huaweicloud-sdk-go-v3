package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInstancesRequest Request Object
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
