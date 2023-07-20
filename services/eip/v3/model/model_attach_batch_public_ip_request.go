package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachBatchPublicIpRequest Request Object
type AttachBatchPublicIpRequest struct {
	Body *BatchAttachSharebwReq `json:"body,omitempty"`
}

func (o AttachBatchPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachBatchPublicIpRequest struct{}"
	}

	return strings.Join([]string{"AttachBatchPublicIpRequest", string(data)}, " ")
}
