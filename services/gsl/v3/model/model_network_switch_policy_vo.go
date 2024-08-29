package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkSwitchPolicyVo struct {

	// 策略主键id
	Id *int64 `json:"id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 最优选运营商,1:移动、2:电信、3:联通、4:上次使用的运营商
	PreferredCarrier *int32 `json:"preferred_carrier,omitempty"`

	// 最不优选运营商,1:移动、2:电信、3:联通
	LeastPreferredCarrier *int32 `json:"least_preferred_carrier,omitempty"`

	// 最优信号选取策略是否开启,true:开启,false:不开启
	OptimalSignal *bool `json:"optimal_signal,omitempty"`

	// 自动切卡是否开启,true:开启,false:不开启
	AutoSwitch *bool `json:"auto_switch,omitempty"`

	// 弱信号切换策略是否开启,true:开启,false:不开启
	WeakSignalSwitch *bool `json:"weak_signal_switch,omitempty"`

	// 连接延时切换策略，连接延时时需要ping的ip地址
	ConnectIp *string `json:"connect_ip,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 修改时间
	ModifyTime *sdktime.SdkTime `json:"modify_time,omitempty"`

	// 状态,1:启用、0:禁用
	Status *int32 `json:"status,omitempty"`

	// 策略类型,0:系统策略,1:私有策略
	Type *int32 `json:"type,omitempty"`

	// 策略是否已被使用过，即是否已被执行过策略切换，false：未被使用过，true：已被使用过
	Used *bool `json:"used,omitempty"`

	// 版本枚举1-SDK版 2-无SDK版
	Version *int32 `json:"version,omitempty"`

	// 切卡顺序，运营商以英文逗号分隔
	SwitchOrder *string `json:"switch_order,omitempty"`

	// 黑名单，只支持单个运营商
	Blacklist *int32 `json:"blacklist,omitempty"`
}

func (o NetworkSwitchPolicyVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkSwitchPolicyVo struct{}"
	}

	return strings.Join([]string{"NetworkSwitchPolicyVo", string(data)}, " ")
}
