package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtraInfo **参数解释** 指标信息 **约束限制** 不涉及
type ExtraInfo struct {

	// **参数解释** 指标名称 **约束限制** 不涉及 **取值范围** 长度为[1,4096]个字符 **默认取值** 不涉及
	OriginMetricName string `json:"origin_metric_name"`

	// **参数解释** 指标名称前缀 **约束限制** 不涉及           **取值范围** 长度为[1,4096]个字符 **默认取值** 不涉及
	MetricPrefix *string `json:"metric_prefix,omitempty"`

	// **参数解释** 指标类型 **约束限制** 不涉及           **取值范围** 长度为[1,32]个字符 **默认取值** 不涉及
	MetricType *string `json:"metric_type,omitempty"`

	// **参数解释** 自定义进程名称 **约束限制** 不涉及           **取值范围** 长度为[1,250]个字符 **默认取值** 不涉及
	CustomProcName *string `json:"custom_proc_name,omitempty"`
}

func (o ExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtraInfo struct{}"
	}

	return strings.Join([]string{"ExtraInfo", string(data)}, " ")
}
