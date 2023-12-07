package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNatGatewayTagResponse Response Object
type CreateNatGatewayTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateNatGatewayTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayTagResponse struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayTagResponse", string(data)}, " ")
}
