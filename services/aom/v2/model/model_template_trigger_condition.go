package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TemplateTriggerCondition struct {

	// 指标查询模式。 - “AOM”：AOM原生 - “PROM”：AOM Prometheus - “NATIVE_PROM”：原生Prometheus
	MetricQueryMode *TemplateTriggerConditionMetricQueryMode `json:"metric_query_mode,omitempty"`

	// 指标命名空间。
	MetricNamespace *string `json:"metric_namespace,omitempty"`

	// 指标名称。
	MetricName *string `json:"metric_name,omitempty"`

	// 指标单位。
	MetricUnit *string `json:"metric_unit,omitempty"`

	// 指标维度。
	MetricLabels *[]string `json:"metric_labels,omitempty"`

	// Prometheus语句。
	Promql *string `json:"promql,omitempty"`

	// Prometheus语句模板。
	PromqlExpr *string `json:"promql_expr,omitempty"`

	// 连续周期个数。
	TriggerTimes *int32 `json:"trigger_times,omitempty"`

	// 检查频率周期。 - 当trigger_type 为“HOURLY”时，填“” - 当trigger_type为“DAILY”时，格式为：“小时” 。例如每天凌晨三点\"03:00\"。 - 当trigger_type为“WEEKLY”时，格式为：“星期 小时”。例如每周一凌晨三点 “1 03:00”。 - 当trigger_type为“CRON”时，格式为标准CRON表达式。 - 当trigger_type为“FIXED_RATE”时，秒的取值为15s，30s，分钟为 1~59，小时为 1~24。例如“15s”，“30s”，“1min”，“1h”。
	TriggerInterval *string `json:"trigger_interval,omitempty"`

	// 触发频率的类型： - “FIXED_RATE”：固定间隔 - “HOURLY”：每小时 - “DAILY”：每天 - “WEEKLY”：每周 - “CRON”：Cron表达式
	TriggerType *TemplateTriggerConditionTriggerType `json:"trigger_type,omitempty"`

	// 持续时间。 持续时间范围： - “0”：立即 - “15s”：15秒 - “30s”：30秒 - “1m”：1分钟
	PromqlFor *string `json:"promql_for,omitempty"`

	// 检测规则： - average：平均值 - minimum：最小值 - maximum：最大值 - sum：总计 - sampleCount：样本
	AggregationType *string `json:"aggregation_type,omitempty"`

	// 判断条件：“>”,“<”,“=”,“>=”,“<=”
	Operator *string `json:"operator,omitempty"`

	// 键值对形式，键为告警级别，值为告警阈值
	Thresholds map[string]string `json:"thresholds,omitempty"`

	// 统计周期。 - “15s”：15秒 - “30s”：30秒 - “1m”：1分钟 - “5m”：5分钟 - “15m”：15分钟 - “1h”：1小时
	AggregationWindow *string `json:"aggregation_window,omitempty"`

	Cmdb *CmdbInfo `json:"cmdb,omitempty"`

	// 查询筛选条件。
	QueryMatch *string `json:"query_match,omitempty"`

	// 监控层级。
	AomMonitorLevel *string `json:"aom_monitor_level,omitempty"`

	// 聚合方式。 - “by”：不分组 - “avg” - “max” - “min” - “sum”
	AggregateType *TemplateTriggerConditionAggregateType `json:"aggregate_type,omitempty"`

	// 当配置方式为全量指标时可选择的指标运算方式。 - “single”：单个指标进行运算 - “mix”：多个指标进行混合运算
	MetricStatisticMethod *TemplateTriggerConditionMetricStatisticMethod `json:"metric_statistic_method,omitempty"`

	// 混合运算的表达式。
	Expression *string `json:"expression,omitempty"`

	// 混合运算的promQL。
	MixPromql *string `json:"mix_promql,omitempty"`

	// 通知内容。
	AlarmMessageTemplate *string `json:"alarm_message_template,omitempty"`

	// Prometheus监控模板。默认为cluster。
	PromqlMonitorTemplates *[]string `json:"promql_monitor_templates,omitempty"`

	// - PromQL告警规则，该参数为\"{\\\"defaultRule\\\":{\\\"label\\\":\\\"自定义\\\",\\\"id\\\":\\\"custom\\\"},\\\"templateSelectd\\\":null,\\\"dimensionsList\\\":[]}\" - 阈值告警规则，该参数为空。
	QueryParam *interface{} `json:"query_param,omitempty"`
}

func (o TemplateTriggerCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateTriggerCondition struct{}"
	}

	return strings.Join([]string{"TemplateTriggerCondition", string(data)}, " ")
}

type TemplateTriggerConditionMetricQueryMode struct {
	value string
}

type TemplateTriggerConditionMetricQueryModeEnum struct {
	AOM         TemplateTriggerConditionMetricQueryMode
	PROM        TemplateTriggerConditionMetricQueryMode
	NATIVE_PROM TemplateTriggerConditionMetricQueryMode
}

func GetTemplateTriggerConditionMetricQueryModeEnum() TemplateTriggerConditionMetricQueryModeEnum {
	return TemplateTriggerConditionMetricQueryModeEnum{
		AOM: TemplateTriggerConditionMetricQueryMode{
			value: "AOM",
		},
		PROM: TemplateTriggerConditionMetricQueryMode{
			value: "PROM",
		},
		NATIVE_PROM: TemplateTriggerConditionMetricQueryMode{
			value: "NATIVE_PROM",
		},
	}
}

func (c TemplateTriggerConditionMetricQueryMode) Value() string {
	return c.value
}

func (c TemplateTriggerConditionMetricQueryMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateTriggerConditionMetricQueryMode) UnmarshalJSON(b []byte) error {
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

type TemplateTriggerConditionTriggerType struct {
	value string
}

type TemplateTriggerConditionTriggerTypeEnum struct {
	FIXED_RATE TemplateTriggerConditionTriggerType
	HOURLY     TemplateTriggerConditionTriggerType
	DAILY      TemplateTriggerConditionTriggerType
	WEEKLY     TemplateTriggerConditionTriggerType
	CRON       TemplateTriggerConditionTriggerType
}

func GetTemplateTriggerConditionTriggerTypeEnum() TemplateTriggerConditionTriggerTypeEnum {
	return TemplateTriggerConditionTriggerTypeEnum{
		FIXED_RATE: TemplateTriggerConditionTriggerType{
			value: "FIXED_RATE",
		},
		HOURLY: TemplateTriggerConditionTriggerType{
			value: "HOURLY",
		},
		DAILY: TemplateTriggerConditionTriggerType{
			value: "DAILY",
		},
		WEEKLY: TemplateTriggerConditionTriggerType{
			value: "WEEKLY",
		},
		CRON: TemplateTriggerConditionTriggerType{
			value: "CRON",
		},
	}
}

func (c TemplateTriggerConditionTriggerType) Value() string {
	return c.value
}

func (c TemplateTriggerConditionTriggerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateTriggerConditionTriggerType) UnmarshalJSON(b []byte) error {
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

type TemplateTriggerConditionAggregateType struct {
	value string
}

type TemplateTriggerConditionAggregateTypeEnum struct {
	BY  TemplateTriggerConditionAggregateType
	AVG TemplateTriggerConditionAggregateType
	MAX TemplateTriggerConditionAggregateType
	MIN TemplateTriggerConditionAggregateType
	SUM TemplateTriggerConditionAggregateType
}

func GetTemplateTriggerConditionAggregateTypeEnum() TemplateTriggerConditionAggregateTypeEnum {
	return TemplateTriggerConditionAggregateTypeEnum{
		BY: TemplateTriggerConditionAggregateType{
			value: "by",
		},
		AVG: TemplateTriggerConditionAggregateType{
			value: "avg",
		},
		MAX: TemplateTriggerConditionAggregateType{
			value: "max",
		},
		MIN: TemplateTriggerConditionAggregateType{
			value: "min",
		},
		SUM: TemplateTriggerConditionAggregateType{
			value: "sum",
		},
	}
}

func (c TemplateTriggerConditionAggregateType) Value() string {
	return c.value
}

func (c TemplateTriggerConditionAggregateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateTriggerConditionAggregateType) UnmarshalJSON(b []byte) error {
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

type TemplateTriggerConditionMetricStatisticMethod struct {
	value string
}

type TemplateTriggerConditionMetricStatisticMethodEnum struct {
	SINGLE TemplateTriggerConditionMetricStatisticMethod
	MIX    TemplateTriggerConditionMetricStatisticMethod
}

func GetTemplateTriggerConditionMetricStatisticMethodEnum() TemplateTriggerConditionMetricStatisticMethodEnum {
	return TemplateTriggerConditionMetricStatisticMethodEnum{
		SINGLE: TemplateTriggerConditionMetricStatisticMethod{
			value: "single",
		},
		MIX: TemplateTriggerConditionMetricStatisticMethod{
			value: "mix",
		},
	}
}

func (c TemplateTriggerConditionMetricStatisticMethod) Value() string {
	return c.value
}

func (c TemplateTriggerConditionMetricStatisticMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateTriggerConditionMetricStatisticMethod) UnmarshalJSON(b []byte) error {
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
