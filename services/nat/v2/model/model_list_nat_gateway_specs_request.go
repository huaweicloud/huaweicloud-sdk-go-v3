package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNatGatewaySpecsRequest Request Object
type ListNatGatewaySpecsRequest struct {
}

func (o ListNatGatewaySpecsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewaySpecsRequest struct{}"
	}

	return strings.Join([]string{"ListNatGatewaySpecsRequest", string(data)}, " ")
}
