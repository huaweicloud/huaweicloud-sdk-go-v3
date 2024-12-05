package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRmsGlobalDcGatewayResponse Response Object
type ListRmsGlobalDcGatewayResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 全球接入网关
	GlobalDcGateways *[]RmsListGlobalDcGateways `json:"global_dc_gateways,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRmsGlobalDcGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRmsGlobalDcGatewayResponse struct{}"
	}

	return strings.Join([]string{"ListRmsGlobalDcGatewayResponse", string(data)}, " ")
}
