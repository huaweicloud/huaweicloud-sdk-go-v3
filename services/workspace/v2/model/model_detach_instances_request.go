package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachInstancesRequest Request Object
type DetachInstancesRequest struct {
	Body *DetachInstancesReq `json:"body,omitempty"`
}

func (o DetachInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachInstancesRequest struct{}"
	}

	return strings.Join([]string{"DetachInstancesRequest", string(data)}, " ")
}
