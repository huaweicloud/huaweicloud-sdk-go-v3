package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAttachInstancesRequest Request Object
type BatchAttachInstancesRequest struct {
	Body *BatchAttachInstancesReq `json:"body,omitempty"`
}

func (o BatchAttachInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchAttachInstancesRequest", string(data)}, " ")
}
