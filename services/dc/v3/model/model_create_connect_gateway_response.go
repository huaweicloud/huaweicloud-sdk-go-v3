package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectGatewayResponse Response Object
type CreateConnectGatewayResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	ConnectGateway *ConnectGatewayResponse `json:"connect_gateway,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreateConnectGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectGatewayResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectGatewayResponse", string(data)}, " ")
}
