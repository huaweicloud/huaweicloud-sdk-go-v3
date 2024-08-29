package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimDeviceMultiplyVo struct {

	// sim卡id
	SimCardId *int64 `json:"sim_card_id,omitempty"`

	// 三网卡标识
	Cid *string `json:"cid,omitempty"`

	// 在线运营商标识
	OnlineCarrier *int32 `json:"online_carrier,omitempty"`

	// 网络切换策略id
	NetworkSwitchPolicyId *int64 `json:"network_switch_policy_id,omitempty"`

	// 网络切换策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 版本信息
	Version *int32 `json:"version,omitempty"`
}

func (o SimDeviceMultiplyVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimDeviceMultiplyVo struct{}"
	}

	return strings.Join([]string{"SimDeviceMultiplyVo", string(data)}, " ")
}
