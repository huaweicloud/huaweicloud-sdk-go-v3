package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalDcGatewayResponse Response Object
type ShowGlobalDcGatewayResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	GlobalDcGateway *ExternalShowGlobalDcGateway `json:"global_dc_gateway,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowGlobalDcGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalDcGatewayResponse struct{}"
	}

	return strings.Join([]string{"ShowGlobalDcGatewayResponse", string(data)}, " ")
}
