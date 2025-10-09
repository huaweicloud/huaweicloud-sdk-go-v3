package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeP2cVpnGatewayResponse Response Object
type UpgradeP2cVpnGatewayResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeP2cVpnGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeP2cVpnGatewayResponse struct{}"
	}

	return strings.Join([]string{"UpgradeP2cVpnGatewayResponse", string(data)}, " ")
}
