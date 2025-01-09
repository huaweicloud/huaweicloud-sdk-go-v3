package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDetachInstancesRequest Request Object
type BatchDetachInstancesRequest struct {
	Body *BatchDetachInstancesReq `json:"body,omitempty"`
}

func (o BatchDetachInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDetachInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchDetachInstancesRequest", string(data)}, " ")
}
