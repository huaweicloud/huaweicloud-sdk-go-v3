package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricExtraInfoResp **参数解释**： 告警策略附加信息。
type MetricExtraInfoResp struct {

	// **参数解释**： 原始指标名称。 **取值范围**： 长度为[0,4096]个字符。
	OriginMetricName *string `json:"origin_metric_name,omitempty"`

	// **参数解释**： 指标名称前缀。 **取值范围**： 长度为[0,4096]个字符。
	MetricPrefix *string `json:"metric_prefix,omitempty"`

	// **参数解释**： 用户进程名称。 **取值范围**： 长度为[0,250]个字符。
	CustomProcName *string `json:"custom_proc_name,omitempty"`

	// **参数解释**： 指标类型。 **取值范围**： 长度为[0,32]个字符。
	MetricType *string `json:"metric_type,omitempty"`
}

func (o MetricExtraInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricExtraInfoResp struct{}"
	}

	return strings.Join([]string{"MetricExtraInfoResp", string(data)}, " ")
}
