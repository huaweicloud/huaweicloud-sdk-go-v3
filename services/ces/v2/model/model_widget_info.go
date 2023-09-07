package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WidgetInfo struct {

	// 指标列表
	Metrics *[]WidgetMetric `json:"metrics,omitempty"`

	// 监控视图标题
	Title *string `json:"title,omitempty"`

	// 监控视图指标的阈值
	Threshold *float64 `json:"threshold,omitempty"`

	// 阈值是否展示，true:展示，false:不展示
	ThresholdEnabled *bool `json:"threshold_enabled,omitempty"`

	// 监控视图图表类型, bar柱状图，line折线图
	View *WidgetInfoView `json:"view,omitempty"`

	// 指标展示类型，single 单指标展示，multiple 多指标展示
	MetricDisplayMode *WidgetInfoMetricDisplayMode `json:"metric_display_mode,omitempty"`

	Properties *BaseWidgetInfoProperties `json:"properties,omitempty"`

	Location *BaseWidgetInfoLocation `json:"location,omitempty"`

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
	BAR  WidgetInfoView
	LINE WidgetInfoView
}

func GetWidgetInfoViewEnum() WidgetInfoViewEnum {
	return WidgetInfoViewEnum{
		BAR: WidgetInfoView{
			value: "bar",
		},
		LINE: WidgetInfoView{
			value: "line",
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
