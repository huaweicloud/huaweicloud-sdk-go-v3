package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MetricMetaDataDerivedMetrics struct {

	// 衍生指标结果维度，0维：单个数字，2维：图表或表格，3+维：多标签图表
	MetricDimension int32 `json:"metric_dimension"`

	// 指标支持的最大检索范围，单位：天；
	MaxQueryRange *int32 `json:"max_query_range,omitempty"`

	// 指标查询范围相对起始时间 datemath表达式
	DateStart *string `json:"date_start,omitempty"`

	// 指标查询范围相对截止时间 datemath表达式
	DateEnd *string `json:"date_end,omitempty"`

	// 时间格式，epoch_millis;epoch_second;yyyy-MM-dd'T'HH:mm:ss.SSSZ
	DateFormat *string `json:"date_format,omitempty"`

	// 获取指标结果方式，cbsl, api, dsl, sql
	QueryType MetricMetaDataDerivedMetricsQueryType `json:"query_type"`

	QueryFunction string `json:"query_function"`
}

func (o MetricMetaDataDerivedMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricMetaDataDerivedMetrics struct{}"
	}

	return strings.Join([]string{"MetricMetaDataDerivedMetrics", string(data)}, " ")
}

type MetricMetaDataDerivedMetricsQueryType struct {
	value string
}

type MetricMetaDataDerivedMetricsQueryTypeEnum struct {
	CBSL MetricMetaDataDerivedMetricsQueryType
	API  MetricMetaDataDerivedMetricsQueryType
	DSL  MetricMetaDataDerivedMetricsQueryType
	SQL  MetricMetaDataDerivedMetricsQueryType
}

func GetMetricMetaDataDerivedMetricsQueryTypeEnum() MetricMetaDataDerivedMetricsQueryTypeEnum {
	return MetricMetaDataDerivedMetricsQueryTypeEnum{
		CBSL: MetricMetaDataDerivedMetricsQueryType{
			value: "cbsl",
		},
		API: MetricMetaDataDerivedMetricsQueryType{
			value: " api",
		},
		DSL: MetricMetaDataDerivedMetricsQueryType{
			value: " dsl",
		},
		SQL: MetricMetaDataDerivedMetricsQueryType{
			value: " sql",
		},
	}
}

func (c MetricMetaDataDerivedMetricsQueryType) Value() string {
	return c.value
}

func (c MetricMetaDataDerivedMetricsQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MetricMetaDataDerivedMetricsQueryType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
