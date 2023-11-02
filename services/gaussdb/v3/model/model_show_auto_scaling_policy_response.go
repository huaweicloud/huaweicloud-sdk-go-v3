package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoScalingPolicyResponse Response Object
type ShowAutoScalingPolicyResponse struct {

	// 自动变配策略ID。
	Id *string `json:"id,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 自动变配开关状态。  取值：  - ON：已开启 - OFF：已关闭
	Status *string `json:"status,omitempty"`

	// 监测周期（单位：秒）。 在整个观测窗口期内，若CPU平均使用率大于等于设定值，则在观测窗口结束后，进行扩容。
	MonitorCycle *int32 `json:"monitor_cycle,omitempty"`

	// 静默周期（单位：秒）。 两次自动扩容或自动回缩的最小间隔时间。
	SilenceCycle *int32 `json:"silence_cycle,omitempty"`

	// 扩容阈值（百分比数值），指CPU平均使用率。
	EnlargeThreshold *int32 `json:"enlarge_threshold,omitempty"`

	// 扩容规格上限。
	MaxFlavor *string `json:"max_flavor,omitempty"`

	// 自动回缩开关状态。  取值：  - true：已开启 - false：已关闭
	ReduceEnabled *bool `json:"reduce_enabled,omitempty"`

	// 缩容规格下限。
	MinFlavor *string `json:"min_flavor,omitempty"`

	// 静默期开始时间（上一次变更结束时间）。  格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	SilenceStartAt *string `json:"silence_start_at,omitempty"`

	ScalingStrategy *ScalingStrategyInfo `json:"scaling_strategy,omitempty"`

	// 只读节点数量上限。
	MaxReadOnlyCount *int32 `json:"max_read_only_count,omitempty"`

	// 只读节点数量下限。
	MinReadOnlyCount *int32 `json:"min_read_only_count,omitempty"`

	// 只读节点读写分离权重。
	ReadOnlyWeight *int32 `json:"read_only_weight,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAutoScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoScalingPolicyResponse", string(data)}, " ")
}
