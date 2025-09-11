package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BaseWidgetInfo **参数解释** 监控视图信息 **约束限制** 不涉及
type BaseWidgetInfo struct {

	// **参数解释** 视图分组id **约束限制** 不涉及           **取值范围** 字符串必须以dg开头，包含22个字母和数字，长度为24个字符 **默认取值** 不涉及
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释** 指标列表 **约束限制** 包含的指标数量最多为200个，最少为1个
	Metrics []WidgetMetric `json:"metrics"`

	// **参数解释** 监控视图标题 **约束限制** 不涉及           **取值范围** 长度为[1,128]个字符，允许包括以下内容：1、中文汉字；2、拉丁字母；3、英文大小写字母；4、数字(0-9)；5、符号： ” \" ≤ < > & % _ : / ; “ ' ? + , ~ ， （ ） º ( ) [ . - **默认取值** 不涉及
	Title string `json:"title"`

	// **参数解释** 监控视图指标的阈值 **约束限制** 不涉及   **取值范围** 最小值为0，最大值为1.7976931348623157e+308 **默认取值** 不涉及
	Threshold *float64 `json:"threshold,omitempty"`

	// **参数解释** 阈值是否展示 **约束限制** 不涉及   **取值范围** - true:展示 - false:不展示 **默认取值** 不涉及
	ThresholdEnabled bool `json:"threshold_enabled"`

	// **参数解释** 监控视图图表类型 **约束限制** 不涉及   **取值范围** - bar:条形图 - line:折线图 - bar_chart:柱状图 - table:表格 - circular_bar:环形柱状图 - area_chart:面积图 **默认取值** 不涉及
	View BaseWidgetInfoView `json:"view"`

	// **参数解释** 指标展示类型 **约束限制** 不涉及   **取值范围** - single:单指标展示 - multiple:多指标展示 **默认取值** 不涉及
	MetricDisplayMode BaseWidgetInfoMetricDisplayMode `json:"metric_display_mode"`

	Properties *BaseWidgetInfoProperties `json:"properties,omitempty"`

	Location *BaseWidgetInfoLocation `json:"location"`

	// **参数解释** 单位 **约束限制** 不涉及 **取值范围** 长度为[0,32]个字符 **默认取值** 不涉及
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
	BAR          BaseWidgetInfoView
	LINE         BaseWidgetInfoView
	BAR_CHART    BaseWidgetInfoView
	TABLE        BaseWidgetInfoView
	CIRCULAR_BAR BaseWidgetInfoView
	AREA_CHART   BaseWidgetInfoView
}

func GetBaseWidgetInfoViewEnum() BaseWidgetInfoViewEnum {
	return BaseWidgetInfoViewEnum{
		BAR: BaseWidgetInfoView{
			value: "bar",
		},
		LINE: BaseWidgetInfoView{
			value: "line",
		},
		BAR_CHART: BaseWidgetInfoView{
			value: "bar_chart",
		},
		TABLE: BaseWidgetInfoView{
			value: "table",
		},
		CIRCULAR_BAR: BaseWidgetInfoView{
			value: "circular_bar",
		},
		AREA_CHART: BaseWidgetInfoView{
			value: "area_chart",
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
