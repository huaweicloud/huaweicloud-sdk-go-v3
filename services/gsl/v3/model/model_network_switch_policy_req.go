package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkSwitchPolicyReq struct {

	// 容器ID
	Cid *string `json:"cid,omitempty"`

	// 网络切换策略标识
	NetworkSwitchPolicyId int64 `json:"network_switch_policy_id"`
}

func (o NetworkSwitchPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkSwitchPolicyReq struct{}"
	}

	return strings.Join([]string{"NetworkSwitchPolicyReq", string(data)}, " ")
}
