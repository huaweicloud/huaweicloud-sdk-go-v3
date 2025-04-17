package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WidgetInfoWithId struct {

	// 视图id
	WidgetId *string `json:"widget_id,omitempty"`

	// 视图分区id
	GroupId *string `json:"group_id,omitempty"`

	// 指标列表
	Metrics []WidgetMetric `json:"metrics"`

	// 监控视图标题
	Title string `json:"title"`

	// 监控视图指标的阈值
	Threshold *float64 `json:"threshold,omitempty"`

	// 阈值是否展示，true:展示，false:不展示
	ThresholdEnabled bool `json:"threshold_enabled"`

	// 监控视图图表类型, bar条形图，line折线图，bar_chart柱状图，table表格，circular_bar环形柱状图，area_chart面积图
	View WidgetInfoWithIdView `json:"view"`

	// 指标展示类型，single 单指标展示，multiple 多指标展示
	MetricDisplayMode WidgetInfoWithIdMetricDisplayMode `json:"metric_display_mode"`

	Properties *BaseWidgetInfoProperties `json:"properties,omitempty"`

	Location *UpdateWidgetInfoLocation `json:"location"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 监控看板创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o WidgetInfoWithId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WidgetInfoWithId struct{}"
	}

	return strings.Join([]string{"WidgetInfoWithId", string(data)}, " ")
}

type WidgetInfoWithIdView struct {
	value string
}

type WidgetInfoWithIdViewEnum struct {
	BAR          WidgetInfoWithIdView
	LINE         WidgetInfoWithIdView
	BAR_CHART    WidgetInfoWithIdView
	TABLE        WidgetInfoWithIdView
	CIRCULAR_BAR WidgetInfoWithIdView
	AREA_CHART   WidgetInfoWithIdView
}

func GetWidgetInfoWithIdViewEnum() WidgetInfoWithIdViewEnum {
	return WidgetInfoWithIdViewEnum{
		BAR: WidgetInfoWithIdView{
			value: "bar",
		},
		LINE: WidgetInfoWithIdView{
			value: "line",
		},
		BAR_CHART: WidgetInfoWithIdView{
			value: "bar_chart",
		},
		TABLE: WidgetInfoWithIdView{
			value: "table",
		},
		CIRCULAR_BAR: WidgetInfoWithIdView{
			value: "circular_bar",
		},
		AREA_CHART: WidgetInfoWithIdView{
			value: "area_chart",
		},
	}
}

func (c WidgetInfoWithIdView) Value() string {
	return c.value
}

func (c WidgetInfoWithIdView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WidgetInfoWithIdView) UnmarshalJSON(b []byte) error {
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

type WidgetInfoWithIdMetricDisplayMode struct {
	value string
}

type WidgetInfoWithIdMetricDisplayModeEnum struct {
	SINGLE   WidgetInfoWithIdMetricDisplayMode
	MULTIPLE WidgetInfoWithIdMetricDisplayMode
}

func GetWidgetInfoWithIdMetricDisplayModeEnum() WidgetInfoWithIdMetricDisplayModeEnum {
	return WidgetInfoWithIdMetricDisplayModeEnum{
		SINGLE: WidgetInfoWithIdMetricDisplayMode{
			value: "single",
		},
		MULTIPLE: WidgetInfoWithIdMetricDisplayMode{
			value: "multiple",
		},
	}
}

func (c WidgetInfoWithIdMetricDisplayMode) Value() string {
	return c.value
}

func (c WidgetInfoWithIdMetricDisplayMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WidgetInfoWithIdMetricDisplayMode) UnmarshalJSON(b []byte) error {
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
