package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVirtualGatewayRequest struct {
	Body *CreateVirtualGatewayRequestBody `json:"body,omitempty"`
}

func (o CreateVirtualGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualGatewayRequest struct{}"
	}

	return strings.Join([]string{"CreateVirtualGatewayRequest", string(data)}, " ")
}
