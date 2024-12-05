package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalDcGatewayRequest Request Object
type DeleteGlobalDcGatewayRequest struct {

	// 全球接入网关ID
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`
}

func (o DeleteGlobalDcGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalDcGatewayRequest struct{}"
	}

	return strings.Join([]string{"DeleteGlobalDcGatewayRequest", string(data)}, " ")
}
