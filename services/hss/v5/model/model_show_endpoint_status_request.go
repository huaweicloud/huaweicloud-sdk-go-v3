package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEndpointStatusRequest Request Object
type ShowEndpointStatusRequest struct {
	Body *QueryVpcCondition `json:"body,omitempty"`
}

func (o ShowEndpointStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEndpointStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowEndpointStatusRequest", string(data)}, " ")
}
