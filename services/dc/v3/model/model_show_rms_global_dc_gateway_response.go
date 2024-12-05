package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRmsGlobalDcGatewayResponse Response Object
type ShowRmsGlobalDcGatewayResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	GlobalDcGateway *RmsShowGlobalDcGateway `json:"global_dc_gateway,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRmsGlobalDcGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRmsGlobalDcGatewayResponse struct{}"
	}

	return strings.Join([]string{"ShowRmsGlobalDcGatewayResponse", string(data)}, " ")
}
