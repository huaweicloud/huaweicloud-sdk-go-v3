package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectGatewayResponse Response Object
type ShowConnectGatewayResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	ConnectGateway *ConnectGatewayResponse `json:"connect_gateway,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowConnectGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectGatewayResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectGatewayResponse", string(data)}, " ")
}
