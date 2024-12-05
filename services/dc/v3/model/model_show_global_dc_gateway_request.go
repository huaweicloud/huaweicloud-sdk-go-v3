package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalDcGatewayRequest Request Object
type ShowGlobalDcGatewayRequest struct {

	// 全球接入网关ID
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// show response ext-fields
	ExtFields *[]string `json:"ext_fields,omitempty"`
}

func (o ShowGlobalDcGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalDcGatewayRequest struct{}"
	}

	return strings.Join([]string{"ShowGlobalDcGatewayRequest", string(data)}, " ")
}
