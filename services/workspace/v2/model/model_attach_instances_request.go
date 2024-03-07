package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachInstancesRequest Request Object
type AttachInstancesRequest struct {
	Body *AttachInstancesReq `json:"body,omitempty"`
}

func (o AttachInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInstancesRequest struct{}"
	}

	return strings.Join([]string{"AttachInstancesRequest", string(data)}, " ")
}
