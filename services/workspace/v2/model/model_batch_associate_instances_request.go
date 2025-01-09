package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAssociateInstancesRequest Request Object
type BatchAssociateInstancesRequest struct {
	Body *PreBatchAttachInstancesReq `json:"body,omitempty"`
}

func (o BatchAssociateInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchAssociateInstancesRequest", string(data)}, " ")
}
