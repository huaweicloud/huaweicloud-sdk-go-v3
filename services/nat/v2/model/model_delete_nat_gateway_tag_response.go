package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNatGatewayTagResponse Response Object
type DeleteNatGatewayTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteNatGatewayTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewayTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewayTagResponse", string(data)}, " ")
}
