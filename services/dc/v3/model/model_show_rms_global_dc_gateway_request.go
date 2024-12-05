package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRmsGlobalDcGatewayRequest Request Object
type ShowRmsGlobalDcGatewayRequest struct {

	// 通过rp_name过滤记录
	RpName string `json:"rp_name"`

	// domain_id
	DomainId string `json:"domain_id"`

	// region_id
	RegionId string `json:"region_id"`

	// resource_type
	ResourceType string `json:"resource_type"`

	// 全球接入网关ID
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// show response ext-fields
	ExtFields *[]string `json:"ext_fields,omitempty"`
}

func (o ShowRmsGlobalDcGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRmsGlobalDcGatewayRequest struct{}"
	}

	return strings.Join([]string{"ShowRmsGlobalDcGatewayRequest", string(data)}, " ")
}
