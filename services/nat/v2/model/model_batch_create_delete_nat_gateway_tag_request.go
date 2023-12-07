package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDeleteNatGatewayTagRequest Request Object
type BatchCreateDeleteNatGatewayTagRequest struct {

	// 公网NAT网关ID。
	NatGatewayId string `json:"nat_gateway_id"`

	Body *BatchCreateDeleteNatTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateDeleteNatGatewayTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteNatGatewayTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteNatGatewayTagRequest", string(data)}, " ")
}
