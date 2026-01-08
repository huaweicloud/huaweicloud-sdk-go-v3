package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateWidgetInfo **参数解释** 待修改的监控视图对象 **约束限制** 不涉及
type UpdateWidgetInfo struct {

	// **参数解释** 视图分区id **约束限制** 不涉及 **取值范围** 字符串必须以dg开头，后跟22个字母和数字，总长度为24个字符或者为default，default代表不分组 **默认取值** 不涉及
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释** 视图id **约束限制** 不涉及 **取值范围** 字符串必须以wg开头，后跟22个字母和数字，总长度为24个字符 **默认取值** 不涉及
	WidgetId string `json:"widget_id"`

	// **参数解释** 指标列表 **约束限制** 包含的指标对象个数为[1,200]
	Metrics *[]WidgetMetric `json:"metrics,omitempty"`

	// **参数解释** 监控视图标题 **约束限制** 不涉及 **取值范围** 字符串可以包含中文字符，字母，数字，下划线（_），横线（-），冒号（:），分号（;），左圆括号（(），右圆括号（)），句号（.），波浪线（~）， 中文左括号（（），中文右括号（））长度为[1,128]个字符 **默认取值** 不涉及
	Title *string `json:"title,omitempty"`

	// **参数解释** 监控视图指标的阈值 **约束限制** 不涉及 **取值范围** 阈值为[0,1.7976931348623157e+308] **默认取值** 不涉及
	Threshold *float64 `json:"threshold,omitempty"`

	// **参数解释** 阈值是否展示 **约束限制** 不涉及 **取值范围** - true 展示 - false 不展示 **默认取值** 不涉及
	ThresholdEnabled *bool `json:"threshold_enabled,omitempty"`

	// **参数解释** 监控视图图表类型 **约束限制** 不涉及 **取值范围** 枚举值： - bar 条形图 - line 折线图 - bar_chart 柱状图 - table 表格 - circular_bar 环形柱状图 - area_chart 面积图 **默认取值** 不涉及
	View *UpdateWidgetInfoView `json:"view,omitempty"`

	// **参数解释** 指标展示类型 **约束限制** 不涉及 **取值范围** 枚举值： - single 单指标展示 - multiple 多指标展示 **默认取值** 不涉及
	MetricDisplayMode *UpdateWidgetInfoMetricDisplayMode `json:"metric_display_mode,omitempty"`

	Properties *UpdateWidgetInfoProperties `json:"properties,omitempty"`

	Location *UpdateWidgetInfoLocation `json:"location,omitempty"`

	// **参数解释** 单位 **约束限制** 不涉及 **取值范围** 长度为[0,32]个字符 **默认取值** 不涉及
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
