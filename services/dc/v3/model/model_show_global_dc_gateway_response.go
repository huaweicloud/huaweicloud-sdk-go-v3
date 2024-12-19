package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalDcGatewayResponse Response Object
type ShowGlobalDcGatewayResponse struct {
	GlobalDcGateway *GlobalDcGatewayEntry `json:"global_dc_gateway,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowGlobalDcGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalDcGatewayResponse struct{}"
	}

	return strings.Join([]string{"ShowGlobalDcGatewayResponse", string(data)}, " ")
}
