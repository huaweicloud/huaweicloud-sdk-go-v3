package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectGatewayRequest Request Object
type CreateConnectGatewayRequest struct {
	Body *CreateConnectGatewayRequestBody `json:"body,omitempty"`
}

func (o CreateConnectGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectGatewayRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectGatewayRequest", string(data)}, " ")
}
