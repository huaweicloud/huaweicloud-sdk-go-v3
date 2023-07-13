package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVirtualGatewayResponse Response Object
type DeleteVirtualGatewayResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVirtualGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirtualGatewayResponse struct{}"
	}

	return strings.Join([]string{"DeleteVirtualGatewayResponse", string(data)}, " ")
}
