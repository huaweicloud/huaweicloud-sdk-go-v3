package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalDcGatewaysResponse Response Object
type ListGlobalDcGatewaysResponse struct {

	// 全域接入网关列表。
	GlobalDcGateways *[]GlobalDcGatewayEntry `json:"global_dc_gateways,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGlobalDcGatewaysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalDcGatewaysResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalDcGatewaysResponse", string(data)}, " ")
}
