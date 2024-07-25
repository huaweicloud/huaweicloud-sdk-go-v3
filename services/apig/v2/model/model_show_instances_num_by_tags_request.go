package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstancesNumByTagsRequest Request Object
type ShowInstancesNumByTagsRequest struct {
	Body *TmsQueryReq `json:"body,omitempty"`
}

func (o ShowInstancesNumByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstancesNumByTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowInstancesNumByTagsRequest", string(data)}, " ")
}
