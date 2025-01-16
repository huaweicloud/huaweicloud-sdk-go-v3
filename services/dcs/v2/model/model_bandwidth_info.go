package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BandwidthInfo struct {

	// 临时扩容开始时间
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 临时扩容结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 当前时间
	CurrentTime *int64 `json:"current_time,omitempty"`

	// 当前带宽，单位为GB
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 下一个扩容时间
	NextExpandTime *int64 `json:"next_expand_time,omitempty"`

	// 扩容数量
	ExpandCount *int32 `json:"expand_count,omitempty"`

	// 临时扩容时间间隔
	ExpandEffectTime *int64 `json:"expand_effect_time,omitempty"`

	// 下一次可以扩容间隔时间
	ExpandIntervalTime *int64 `json:"expand_interval_time,omitempty"`

	// 最大扩容数量
	MaxExpandCount *int32 `json:"max_expand_count,omitempty"`

	// 任务是否运行
	TaskRunning *bool `json:"task_running,omitempty"`

	// **参数解释**： 实例基准带宽。 **取值范围**： 不涉及。
	AssuredBandwidth *int32 `json:"assured_bandwidth,omitempty"`

	// **参数解释**： 节点最大带宽。 **取值范围**： 不涉及。
	MaxBandwidthForNode *int32 `json:"max_bandwidth_for_node,omitempty"`
}

func (o BandwidthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthInfo struct{}"
	}

	return strings.Join([]string{"BandwidthInfo", string(data)}, " ")
}
