package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListP2cVpnGatewayJobsRequest Request Object
type ListP2cVpnGatewayJobsRequest struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o ListP2cVpnGatewayJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListP2cVpnGatewayJobsRequest struct{}"
	}

	return strings.Join([]string{"ListP2cVpnGatewayJobsRequest", string(data)}, " ")
}
