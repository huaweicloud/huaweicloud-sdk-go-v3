package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricInfoList **参数解释** 指标信息
type MetricInfoList struct {

	// **参数解释** 指标维度
	Dimensions []MetricsDimensionResp `json:"dimensions"`

	// **参数解释** 指标名称 **取值范围** 不涉及
	MetricName string `json:"metric_name"`

	// **参数解释** 服务命名空间 **取值范围** 不涉及
	Namespace string `json:"namespace"`

	// **参数解释** 指标单位 **取值范围** 不涉及
	Unit string `json:"unit"`
}

func (o MetricInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInfoList struct{}"
	}

	return strings.Join([]string{"MetricInfoList", string(data)}, " ")
}
