package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowWidgetResponse Response Object
type ShowWidgetResponse struct {

	// 视图id
	WidgetId *string `json:"widget_id,omitempty"`

	// 指标列表
	Metrics *[]WidgetMetric `json:"metrics,omitempty"`

	// 监控视图标题
	Title *string `json:"title,omitempty"`

	// 监控视图指标的阈值
	Threshold *float64 `json:"threshold,omitempty"`

	// 阈值是否展示，true:展示，false:不展示
	ThresholdEnabled *bool `json:"threshold_enabled,omitempty"`

	// 监控视图图表类型, bar柱状图，line折线图
	View *ShowWidgetResponseView `json:"view,omitempty"`

	// 指标展示类型，single 单指标展示，multiple 多指标展示
	MetricDisplayMode *ShowWidgetResponseMetricDisplayMode `json:"metric_display_mode,omitempty"`

	Properties *BaseWidgetInfoProperties `json:"properties,omitempty"`

	Location *BaseWidgetInfoLocation `json:"location,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 监控看板创建时间
	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowWidgetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWidgetResponse struct{}"
	}

	return strings.Join([]string{"ShowWidgetResponse", string(data)}, " ")
}

type ShowWidgetResponseView struct {
	value string
}

type ShowWidgetResponseViewEnum struct {
	BAR  ShowWidgetResponseView
	LINE ShowWidgetResponseView
}

func GetShowWidgetResponseViewEnum() ShowWidgetResponseViewEnum {
	return ShowWidgetResponseViewEnum{
		BAR: ShowWidgetResponseView{
			value: "bar",
		},
		LINE: ShowWidgetResponseView{
			value: "line",
		},
	}
}

func (c ShowWidgetResponseView) Value() string {
	return c.value
}

func (c ShowWidgetResponseView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWidgetResponseView) UnmarshalJSON(b []byte) error {
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

type ShowWidgetResponseMetricDisplayMode struct {
	value string
}

type ShowWidgetResponseMetricDisplayModeEnum struct {
	SINGLE   ShowWidgetResponseMetricDisplayMode
	MULTIPLE ShowWidgetResponseMetricDisplayMode
}

func GetShowWidgetResponseMetricDisplayModeEnum() ShowWidgetResponseMetricDisplayModeEnum {
	return ShowWidgetResponseMetricDisplayModeEnum{
		SINGLE: ShowWidgetResponseMetricDisplayMode{
			value: "single",
		},
		MULTIPLE: ShowWidgetResponseMetricDisplayMode{
			value: "multiple",
		},
	}
}

func (c ShowWidgetResponseMetricDisplayMode) Value() string {
	return c.value
}

func (c ShowWidgetResponseMetricDisplayMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWidgetResponseMetricDisplayMode) UnmarshalJSON(b []byte) error {
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
