package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeVpnGatewayResponse Response Object
type UpgradeVpnGatewayResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeVpnGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeVpnGatewayResponse struct{}"
	}

	return strings.Join([]string{"UpgradeVpnGatewayResponse", string(data)}, " ")
}
