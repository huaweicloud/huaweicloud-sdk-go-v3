package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricMetaData 指标元数据
type MetricMetaData struct {

	// 指标名称
	Name string `json:"name"`

	// 指标Id
	Id *string `json:"id,omitempty"`

	// 指标类型， 当前仅支持创建日志类型指标
	MetricType string `json:"metric_type"`

	// 数据类型, 当前仅支持创建统计指标
	DataType string `json:"data_type"`

	// 指标结果维度，0维：单个数字，2维：图表或表格，3+维：多标签图表, metric_type为DERIVED必填，其他类型选填（COMPOUND时必为0）
	MetricDimension *int32 `json:"metric_dimension,omitempty"`

	// 缓存生命周期，单位s
	CacheTtl int64 `json:"cache_ttl"`

	// 上报周期，埋点指标时必填，单位s
	ReportPeriod *int64 `json:"report_period,omitempty"`

	// 是否为系统指标
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 生效的列, 当有该参数时，使用指定列作为指标数据结果
	EffectiveColumn *string `json:"effective_column,omitempty"`

	// 指标支持的最大检索范围，单位：天；复合指标时，数值为derived_metrics列表元素中最小值
	MaxQueryRange *int32 `json:"max_query_range,omitempty"`

	// 衍生指标列表，非复合指标时只有一个元素，复合指标时，为各衍生指标的定义
	DerivedMetrics *[]MetricMetaDataDerivedMetrics `json:"derived_metrics,omitempty"`

	// metric_type为DERIVED时填写, 复合指标的表达式
	CompoundExpression *string `json:"compound_expression,omitempty"`

	// 指标格式
	MetricFormat *[]LayoutMetricFormat `json:"metric_format,omitempty"`

	MetricExpandDim *MetricDimensionExpandParam `json:"metric_expand_dim,omitempty"`

	// 安全云脑版本
	Version *string `json:"version,omitempty"`
}

func (o MetricMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricMetaData struct{}"
	}

	return strings.Join([]string{"MetricMetaData", string(data)}, " ")
}
