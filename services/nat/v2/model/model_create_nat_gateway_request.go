package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateNatGatewayRequest struct {
	Body *CreateNatGatewayRequestBody `json:"body,omitempty"`
}

func (o CreateNatGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayRequest struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayRequest", string(data)}, " ")
}
