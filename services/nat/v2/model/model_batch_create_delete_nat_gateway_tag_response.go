package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDeleteNatGatewayTagResponse Response Object
type BatchCreateDeleteNatGatewayTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateDeleteNatGatewayTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteNatGatewayTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteNatGatewayTagResponse", string(data)}, " ")
}
