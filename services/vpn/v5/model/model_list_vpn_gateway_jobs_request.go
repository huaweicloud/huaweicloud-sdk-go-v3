package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnGatewayJobsRequest Request Object
type ListVpnGatewayJobsRequest struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o ListVpnGatewayJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnGatewayJobsRequest struct{}"
	}

	return strings.Join([]string{"ListVpnGatewayJobsRequest", string(data)}, " ")
}
