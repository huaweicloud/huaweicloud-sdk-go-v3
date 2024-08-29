package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetNetworkSwitchPolicyResponse Response Object
type SetNetworkSwitchPolicyResponse struct {

	// 业务受理单号
	WorkOrderId    *int64 `json:"work_order_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SetNetworkSwitchPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetNetworkSwitchPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetNetworkSwitchPolicyResponse", string(data)}, " ")
}
