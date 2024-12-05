package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalDcGatewayResponse Response Object
type ListGlobalDcGatewayResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	GlobalDcGateways *[]ExternalListGlobalDcGateways `json:"global_dc_gateways,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGlobalDcGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalDcGatewayResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalDcGatewayResponse", string(data)}, " ")
}
