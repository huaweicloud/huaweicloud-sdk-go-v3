package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNatGatewayTagRequest Request Object
type ListNatGatewayTagRequest struct {
}

func (o ListNatGatewayTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewayTagRequest struct{}"
	}

	return strings.Join([]string{"ListNatGatewayTagRequest", string(data)}, " ")
}
