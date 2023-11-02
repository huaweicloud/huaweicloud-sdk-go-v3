package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutoScalingPolicyRequestBody 设置自动变配策略请求体
type UpdateAutoScalingPolicyRequestBody struct {

	// 自动变配开关状态。  取值：  - ON：开启。 - OFF：关闭。
	Status string `json:"status"`

	// 监测周期（单位：秒）。 在整个观测窗口期内，若CPU平均使用率大于等于设定值，则在观测窗口结束后，进行扩容。  取值范围：300、600、900、1800。  status为ON时必填。
	MonitorCycle *int32 `json:"monitor_cycle,omitempty"`

	// 静默周期（单位：秒）。 两次自动扩容或自动回缩的最小间隔时间。  取值范围：300、600、1800、3600、7200、10800、86400、604800。  status为ON时必填。
	SilenceCycle *int32 `json:"silence_cycle,omitempty"`

	// 扩容阈值（百分比数值）。  取值范围：50-100。  status为ON时必填。
	EnlargeThreshold *int32 `json:"enlarge_threshold,omitempty"`

	// 扩容规格上限。开启扩缩规格时必填。
	MaxFlavor *string `json:"max_flavor,omitempty"`

	// 是否开启自动回缩。开启自动变配时必填。 - true：是。 - false：否。
	ReduceEnabled *bool `json:"reduce_enabled,omitempty"`

	// 只读节点数量上限。开启增删只读节点时必填。
	MaxReadOnlyCount *int32 `json:"max_read_only_count,omitempty"`

	// 只读节点读写分离权重。开启增删只读节点时必填。
	ReadOnlyWeight *int32 `json:"read_only_weight,omitempty"`

	ScalingStrategy *ScalingStrategyReqInfo `json:"scaling_strategy"`
}

func (o UpdateAutoScalingPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoScalingPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAutoScalingPolicyRequestBody", string(data)}, " ")
}
