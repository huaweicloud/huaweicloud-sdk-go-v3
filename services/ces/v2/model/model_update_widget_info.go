package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateWidgetInfo 待修改的监控视图对象
type UpdateWidgetInfo struct {

	// 视图分区id
	GroupId *string `json:"group_id,omitempty"`

	// 视图id
	WidgetId string `json:"widget_id"`

	// 指标列表
	Metrics *[]WidgetMetric `json:"metrics,omitempty"`

	// 监控视图标题
	Title *string `json:"title,omitempty"`

	// 监控视图指标的阈值
	Threshold *float64 `json:"threshold,omitempty"`

	// 阈值是否展示，true:展示，false:不展示
	ThresholdEnabled *bool `json:"threshold_enabled,omitempty"`

	// 监控视图图表类型, bar条形图，line折线图，bar_chart柱状图，table表格，circular_bar环形柱状图，area_chart面积图
	View *UpdateWidgetInfoView `json:"view,omitempty"`

	// 指标展示类型，single 单指标展示，multiple 多指标展示
	MetricDisplayMode *UpdateWidgetInfoMetricDisplayMode `json:"metric_display_mode,omitempty"`

	Properties *UpdateWidgetInfoProperties `json:"properties,omitempty"`

	Location *UpdateWidgetInfoLocation `json:"location,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`
}

func (o UpdateWidgetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWidgetInfo struct{}"
	}

	return strings.Join([]string{"UpdateWidgetInfo", string(data)}, " ")
}

type UpdateWidgetInfoView struct {
	value string
}

type UpdateWidgetInfoViewEnum struct {
	BAR          UpdateWidgetInfoView
	LINE         UpdateWidgetInfoView
	BAR_CHART    UpdateWidgetInfoView
	TABLE        UpdateWidgetInfoView
	CIRCULAR_BAR UpdateWidgetInfoView
	AREA_CHART   UpdateWidgetInfoView
}

func GetUpdateWidgetInfoViewEnum() UpdateWidgetInfoViewEnum {
	return UpdateWidgetInfoViewEnum{
		BAR: UpdateWidgetInfoView{
			value: "bar",
		},
		LINE: UpdateWidgetInfoView{
			value: "line",
		},
		BAR_CHART: UpdateWidgetInfoView{
			value: "bar_chart",
		},
		TABLE: UpdateWidgetInfoView{
			value: "table",
		},
		CIRCULAR_BAR: UpdateWidgetInfoView{
			value: "circular_bar",
		},
		AREA_CHART: UpdateWidgetInfoView{
			value: "area_chart",
		},
	}
}

func (c UpdateWidgetInfoView) Value() string {
	return c.value
}

func (c UpdateWidgetInfoView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateWidgetInfoView) UnmarshalJSON(b []byte) error {
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

type UpdateWidgetInfoMetricDisplayMode struct {
	value string
}

type UpdateWidgetInfoMetricDisplayModeEnum struct {
	SINGLE   UpdateWidgetInfoMetricDisplayMode
	MULTIPLE UpdateWidgetInfoMetricDisplayMode
}

func GetUpdateWidgetInfoMetricDisplayModeEnum() UpdateWidgetInfoMetricDisplayModeEnum {
	return UpdateWidgetInfoMetricDisplayModeEnum{
		SINGLE: UpdateWidgetInfoMetricDisplayMode{
			value: "single",
		},
		MULTIPLE: UpdateWidgetInfoMetricDisplayMode{
			value: "multiple",
		},
	}
}

func (c UpdateWidgetInfoMetricDisplayMode) Value() string {
	return c.value
}

func (c UpdateWidgetInfoMetricDisplayMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateWidgetInfoMetricDisplayMode) UnmarshalJSON(b []byte) error {
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
