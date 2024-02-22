package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BaseWidgetInfo 监控视图信息
type BaseWidgetInfo struct {

	// 指标列表
	Metrics *[]WidgetMetric `json:"metrics,omitempty"`

	// 监控视图标题
	Title *string `json:"title,omitempty"`

	// 监控视图指标的阈值
	Threshold *float64 `json:"threshold,omitempty"`

	// 阈值是否展示，true:展示，false:不展示
	ThresholdEnabled *bool `json:"threshold_enabled,omitempty"`

	// 监控视图图表类型, bar柱状图，line折线图
	View *BaseWidgetInfoView `json:"view,omitempty"`

	// 指标展示类型，single 单指标展示，multiple 多指标展示
	MetricDisplayMode *BaseWidgetInfoMetricDisplayMode `json:"metric_display_mode,omitempty"`

	Properties *UpdateWidgetInfoProperties `json:"properties,omitempty"`

	Location *UpdateWidgetInfoLocation `json:"location,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`
}

func (o BaseWidgetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseWidgetInfo struct{}"
	}

	return strings.Join([]string{"BaseWidgetInfo", string(data)}, " ")
}

type BaseWidgetInfoView struct {
	value string
}

type BaseWidgetInfoViewEnum struct {
	BAR  BaseWidgetInfoView
	LINE BaseWidgetInfoView
}

func GetBaseWidgetInfoViewEnum() BaseWidgetInfoViewEnum {
	return BaseWidgetInfoViewEnum{
		BAR: BaseWidgetInfoView{
			value: "bar",
		},
		LINE: BaseWidgetInfoView{
			value: "line",
		},
	}
}

func (c BaseWidgetInfoView) Value() string {
	return c.value
}

func (c BaseWidgetInfoView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoView) UnmarshalJSON(b []byte) error {
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

type BaseWidgetInfoMetricDisplayMode struct {
	value string
}

type BaseWidgetInfoMetricDisplayModeEnum struct {
	SINGLE   BaseWidgetInfoMetricDisplayMode
	MULTIPLE BaseWidgetInfoMetricDisplayMode
}

func GetBaseWidgetInfoMetricDisplayModeEnum() BaseWidgetInfoMetricDisplayModeEnum {
	return BaseWidgetInfoMetricDisplayModeEnum{
		SINGLE: BaseWidgetInfoMetricDisplayMode{
			value: "single",
		},
		MULTIPLE: BaseWidgetInfoMetricDisplayMode{
			value: "multiple",
		},
	}
}

func (c BaseWidgetInfoMetricDisplayMode) Value() string {
	return c.value
}

func (c BaseWidgetInfoMetricDisplayMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoMetricDisplayMode) UnmarshalJSON(b []byte) error {
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
