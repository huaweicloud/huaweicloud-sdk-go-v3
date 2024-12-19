package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalDcGatewayResponse Response Object
type DeleteGlobalDcGatewayResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGlobalDcGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalDcGatewayResponse struct{}"
	}

	return strings.Join([]string{"DeleteGlobalDcGatewayResponse", string(data)}, " ")
}
