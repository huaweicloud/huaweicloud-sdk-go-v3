package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtraInfoResp **参数解释** 指标信息
type ExtraInfoResp struct {

	// **参数解释** 指标名称 **取值范围** 长度为[1,4096]个字符
	OriginMetricName *string `json:"origin_metric_name,omitempty"`

	// **参数解释** 指标名称前缀 **取值范围** 长度为[1,4096]个字符
	MetricPrefix *string `json:"metric_prefix,omitempty"`

	// **参数解释** 指标类型 **取值范围** 长度为[1,32]个字符
	MetricType *string `json:"metric_type,omitempty"`

	// **参数解释** 自定义进程名称 **取值范围** 长度为[1,250]个字符
	CustomProcName *string `json:"custom_proc_name,omitempty"`
}

func (o ExtraInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtraInfoResp struct{}"
	}

	return strings.Join([]string{"ExtraInfoResp", string(data)}, " ")
}
