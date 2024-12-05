package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalDcGatewayResponse Response Object
type CreateGlobalDcGatewayResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	GlobalDcGateway *ExternalCreateGlobalDcGateway `json:"global_dc_gateway,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGlobalDcGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalDcGatewayResponse struct{}"
	}

	return strings.Join([]string{"CreateGlobalDcGatewayResponse", string(data)}, " ")
}
