package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WidgetInfoWithId struct {

	// **参数解释** 视图id **取值范围** 字符串必须以wg开头，包含22个字母和数字，长度为24个字符。
	WidgetId *string `json:"widget_id,omitempty"`

	// **参数解释** 视图分组id **取值范围** 字符串必须以dg开头，包含22个字母和数字，长度为24个字符或者为default，default代表不分组
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释** 指标列表
	Metrics *[]WidgetMetricResp `json:"metrics,omitempty"`

	// **参数解释** 监控视图标题 **取值范围** 长度为[1,128]个字符，允许包括以下内容：1、中文汉字；2、拉丁字母；3、英文大小写字母；4、数字(0-9)；5、符号： ” \" ≤ < > & % _ : / ; “ ' ? + , ~ ， （ ） º ( ) [ . -
	Title *string `json:"title,omitempty"`

	// **参数解释** 监控视图指标的阈值 **取值范围** 最小值为0，最大值为1.7976931348623157e+308
	Threshold *float64 `json:"threshold,omitempty"`

	// **参数解释** 阈值是否展示 **取值范围** - true:展示 - false:不展示
	ThresholdEnabled *bool `json:"threshold_enabled,omitempty"`

	// **参数解释** 监控视图图表类型 **取值范围** - bar:条形图 - line:折线图 - bar_chart:柱状图 - table:表格 - circular_bar:环形柱状图 - area_chart:面积图
	View *WidgetInfoWithIdView `json:"view,omitempty"`

	// **参数解释** 指标展示类型 **取值范围** - single:单指标展示 - multiple:多指标展示
	MetricDisplayMode *WidgetInfoWithIdMetricDisplayMode `json:"metric_display_mode,omitempty"`

	Properties *BaseWidgetInfoRespProperties `json:"properties,omitempty"`

	Location *BaseWidgetInfoRespLocation `json:"location,omitempty"`

	// **参数解释** 单位 **取值范围** 长度为[0,32]个字符
	Unit *string `json:"unit,omitempty"`

	// **参数解释** 监控看板创建时间 **取值范围** 最小值为1111111111111，最大值为9999999999999
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
