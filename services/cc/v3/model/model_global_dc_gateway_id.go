package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlobalDcGatewayId Gdgw的ID。
type GlobalDcGatewayId struct {

	// 资源ID标识符。
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`
}

func (o GlobalDcGatewayId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalDcGatewayId struct{}"
	}

	return strings.Join([]string{"GlobalDcGatewayId", string(data)}, " ")
}
