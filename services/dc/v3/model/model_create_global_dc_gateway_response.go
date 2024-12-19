package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalDcGatewayResponse Response Object
type CreateGlobalDcGatewayResponse struct {
	GlobalDcGateway *CreateGlobalDcGatewayEntry `json:"global_dc_gateway,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGlobalDcGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalDcGatewayResponse struct{}"
	}

	return strings.Join([]string{"CreateGlobalDcGatewayResponse", string(data)}, " ")
}
