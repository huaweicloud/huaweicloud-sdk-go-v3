package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalDcGatewayRequest Request Object
type UpdateGlobalDcGatewayRequest struct {

	// 全球接入网关ID
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	Body *UpdateGlobalDcGatewayRequestBody `json:"body,omitempty"`
}

func (o UpdateGlobalDcGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalDcGatewayRequest struct{}"
	}

	return strings.Join([]string{"UpdateGlobalDcGatewayRequest", string(data)}, " ")
}
