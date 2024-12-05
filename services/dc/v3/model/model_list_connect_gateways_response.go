package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectGatewaysResponse Response Object
type ListConnectGatewaysResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	ConnectGateways *[]ConnectGatewayResponse `json:"connect_gateways,omitempty"`
	HttpStatusCode  int                       `json:"-"`
}

func (o ListConnectGatewaysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectGatewaysResponse struct{}"
	}

	return strings.Join([]string{"ListConnectGatewaysResponse", string(data)}, " ")
}
