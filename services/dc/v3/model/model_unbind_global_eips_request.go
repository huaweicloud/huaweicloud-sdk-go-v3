package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindGlobalEipsRequest Request Object
type UnbindGlobalEipsRequest struct {

	// 互联网关ID
	ConnectGatewayId string `json:"connect_gateway_id"`

	Body *CreateUnbindingGeipRequestBody `json:"body,omitempty"`
}

func (o UnbindGlobalEipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindGlobalEipsRequest struct{}"
	}

	return strings.Join([]string{"UnbindGlobalEipsRequest", string(data)}, " ")
}
