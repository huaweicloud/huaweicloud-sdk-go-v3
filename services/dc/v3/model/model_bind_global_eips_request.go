package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindGlobalEipsRequest Request Object
type BindGlobalEipsRequest struct {

	// 互联网关ID
	ConnectGatewayId string `json:"connect_gateway_id"`

	Body *CreateBindingGeipRequestBody `json:"body,omitempty"`
}

func (o BindGlobalEipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindGlobalEipsRequest struct{}"
	}

	return strings.Join([]string{"BindGlobalEipsRequest", string(data)}, " ")
}
