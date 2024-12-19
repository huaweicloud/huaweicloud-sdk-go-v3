package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalDcGatewayResponse Response Object
type UpdateGlobalDcGatewayResponse struct {
	GlobalDcGateway *GlobalDcGatewayEntry `json:"global_dc_gateway,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGlobalDcGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalDcGatewayResponse struct{}"
	}

	return strings.Join([]string{"UpdateGlobalDcGatewayResponse", string(data)}, " ")
}
