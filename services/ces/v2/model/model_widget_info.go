package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WidgetInfo struct {

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
	View WidgetInfoView `json:"view"`

	// 指标展示类型，single 单指标展示，multiple 多指标展示
	MetricDisplayMode WidgetInfoMetricDisplayMode `json:"metric_display_mode"`

	Properties *BaseWidgetInfoProperties `json:"properties,omitempty"`

	Location *UpdateWidgetInfoLocation `json:"location"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 监控看板创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o WidgetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WidgetInfo struct{}"
	}

	return strings.Join([]string{"WidgetInfo", string(data)}, " ")
}

type WidgetInfoView struct {
	value string
}

type WidgetInfoViewEnum struct {
	BAR          WidgetInfoView
	LINE         WidgetInfoView
	BAR_CHART    WidgetInfoView
	TABLE        WidgetInfoView
	CIRCULAR_BAR WidgetInfoView
	AREA_CHART   WidgetInfoView
}

func GetWidgetInfoViewEnum() WidgetInfoViewEnum {
	return WidgetInfoViewEnum{
		BAR: WidgetInfoView{
			value: "bar",
		},
		LINE: WidgetInfoView{
			value: "line",
		},
		BAR_CHART: WidgetInfoView{
			value: "bar_chart",
		},
		TABLE: WidgetInfoView{
			value: "table",
		},
		CIRCULAR_BAR: WidgetInfoView{
			value: "circular_bar",
		},
		AREA_CHART: WidgetInfoView{
			value: "area_chart",
		},
	}
}

func (c WidgetInfoView) Value() string {
	return c.value
}

func (c WidgetInfoView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WidgetInfoView) UnmarshalJSON(b []byte) error {
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

type WidgetInfoMetricDisplayMode struct {
	value string
}

type WidgetInfoMetricDisplayModeEnum struct {
	SINGLE   WidgetInfoMetricDisplayMode
	MULTIPLE WidgetInfoMetricDisplayMode
}

func GetWidgetInfoMetricDisplayModeEnum() WidgetInfoMetricDisplayModeEnum {
	return WidgetInfoMetricDisplayModeEnum{
		SINGLE: WidgetInfoMetricDisplayMode{
			value: "single",
		},
		MULTIPLE: WidgetInfoMetricDisplayMode{
			value: "multiple",
		},
	}
}

func (c WidgetInfoMetricDisplayMode) Value() string {
	return c.value
}

func (c WidgetInfoMetricDisplayMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WidgetInfoMetricDisplayMode) UnmarshalJSON(b []byte) error {
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
