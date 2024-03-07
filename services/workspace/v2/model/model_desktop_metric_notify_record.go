package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopMetricNotifyRecord 对应指标维度产生的告警记录
type DesktopMetricNotifyRecord struct {

	// 满足通知规则阈值的桌面数
	MatchCount *int32 `json:"match_count,omitempty"`

	// 指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 统计持续周期(天)
	Threshold *int32 `json:"threshold,omitempty"`

	// 统计指标对应的统计值和threshold进行比较的条件 * `>=` -  统计指标大于等于threshold时触发 * `>` -   统计指标大于threshold时触发 * `=` -  统计指标等于threshold时触发 * `<=` -  统计指标小于等于threshold时触发 * `<` -  统计指标小于threshold时触发
	ComparisonOperator *string `json:"comparison_operator,omitempty"`
}

func (o DesktopMetricNotifyRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopMetricNotifyRecord struct{}"
	}

	return strings.Join([]string{"DesktopMetricNotifyRecord", string(data)}, " ")
}
