package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BaseWidgetInfoProperties **参数解释** 额外信息 **约束限制** 不涉及
type BaseWidgetInfoProperties struct {

	// **参数解释** 聚合类型 **约束限制** 折线图不支持该参数 **取值范围** 目前只有TopN这一种类型 **默认取值** 不涉及
	Filter *BaseWidgetInfoPropertiesFilter `json:"filter,omitempty"`

	// **参数解释** Top前N个；折线图时为随机展示的时序数据条数 **约束限制** 不涉及               **取值范围** 最小值为1，最大值为2147483647 **默认取值** 100
	TopN *int32 `json:"topN,omitempty"`

	// **参数解释** 排序字段 **约束限制** 折线图不支持该参数   **取值范围** - asc:正序 - desc:倒序 **默认取值** 不涉及
	Order *BaseWidgetInfoPropertiesOrder `json:"order,omitempty"`

	// **参数解释** 监控视图的描述信息 **约束限制** 不涉及                 **取值范围** 长度为[0,200]个字符 **默认取值** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释** 是否展示同比（上周同一时间）数据 **约束限制** 不涉及                 **取值范围** - true:展示 - false:不展示 **默认取值** 不涉及
	LastWeekCompareEnable *bool `json:"last_week_compare_enable,omitempty"`

	// **参数解释** 是否展示环比（昨天同一时间）数据 **约束限制** 不涉及                 **取值范围** - true:展示 - false:不展示 **默认取值** 不涉及
	YesterdayCompareEnable *bool `json:"yesterday_compare_enable,omitempty"`

	// **参数解释** 图例位置标记 **约束限制** 表格不支持该参数                 **取值范围** - hide:表示隐藏图例 - right:表示图例放在监控视图右侧 - bottom:表示图例放在监控视图底部 **默认取值** 不涉及
	LegendLocation *BaseWidgetInfoPropertiesLegendLocation `json:"legend_location,omitempty"`

	// **参数解释** 当前时序数据需要在图例中展示的统计值名称列表 **约束限制** 表格不支持该参数；条形图和柱状图仅支持配置当前值
	LegendValues *[]BaseWidgetInfoPropertiesLegendValues `json:"legend_values,omitempty"`

	// **参数解释** 监控视图的阈值辅助线配置 **约束限制** 不涉及
	Thresholds *[]ThresholdInfo `json:"thresholds,omitempty"`

	// **参数解释** 同比环比总开关是否生效 **约束限制** 不涉及               **取值范围** - true:生效 - false:不生效 **默认取值** 不涉及
	IsAllCompareEnable *bool `json:"is_all_compare_enable,omitempty"`
}

func (o BaseWidgetInfoProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseWidgetInfoProperties struct{}"
	}

	return strings.Join([]string{"BaseWidgetInfoProperties", string(data)}, " ")
}

type BaseWidgetInfoPropertiesFilter struct {
	value string
}

type BaseWidgetInfoPropertiesFilterEnum struct {
	TOP_N BaseWidgetInfoPropertiesFilter
}

func GetBaseWidgetInfoPropertiesFilterEnum() BaseWidgetInfoPropertiesFilterEnum {
	return BaseWidgetInfoPropertiesFilterEnum{
		TOP_N: BaseWidgetInfoPropertiesFilter{
			value: "topN",
		},
	}
}

func (c BaseWidgetInfoPropertiesFilter) Value() string {
	return c.value
}

func (c BaseWidgetInfoPropertiesFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoPropertiesFilter) UnmarshalJSON(b []byte) error {
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

type BaseWidgetInfoPropertiesOrder struct {
	value string
}

type BaseWidgetInfoPropertiesOrderEnum struct {
	ASC  BaseWidgetInfoPropertiesOrder
	DESC BaseWidgetInfoPropertiesOrder
}

func GetBaseWidgetInfoPropertiesOrderEnum() BaseWidgetInfoPropertiesOrderEnum {
	return BaseWidgetInfoPropertiesOrderEnum{
		ASC: BaseWidgetInfoPropertiesOrder{
			value: "asc",
		},
		DESC: BaseWidgetInfoPropertiesOrder{
			value: "desc",
		},
	}
}

func (c BaseWidgetInfoPropertiesOrder) Value() string {
	return c.value
}

func (c BaseWidgetInfoPropertiesOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoPropertiesOrder) UnmarshalJSON(b []byte) error {
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

type BaseWidgetInfoPropertiesLegendLocation struct {
	value string
}

type BaseWidgetInfoPropertiesLegendLocationEnum struct {
	HIDE   BaseWidgetInfoPropertiesLegendLocation
	RIGHT  BaseWidgetInfoPropertiesLegendLocation
	BOTTOM BaseWidgetInfoPropertiesLegendLocation
}

func GetBaseWidgetInfoPropertiesLegendLocationEnum() BaseWidgetInfoPropertiesLegendLocationEnum {
	return BaseWidgetInfoPropertiesLegendLocationEnum{
		HIDE: BaseWidgetInfoPropertiesLegendLocation{
			value: "hide",
		},
		RIGHT: BaseWidgetInfoPropertiesLegendLocation{
			value: "right",
		},
		BOTTOM: BaseWidgetInfoPropertiesLegendLocation{
			value: "bottom",
		},
	}
}

func (c BaseWidgetInfoPropertiesLegendLocation) Value() string {
	return c.value
}

func (c BaseWidgetInfoPropertiesLegendLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoPropertiesLegendLocation) UnmarshalJSON(b []byte) error {
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

type BaseWidgetInfoPropertiesLegendValues struct {
	value string
}

type BaseWidgetInfoPropertiesLegendValuesEnum struct {
	LAST BaseWidgetInfoPropertiesLegendValues
	MAX  BaseWidgetInfoPropertiesLegendValues
	MIN  BaseWidgetInfoPropertiesLegendValues
	AVG  BaseWidgetInfoPropertiesLegendValues
	SUM  BaseWidgetInfoPropertiesLegendValues
}

func GetBaseWidgetInfoPropertiesLegendValuesEnum() BaseWidgetInfoPropertiesLegendValuesEnum {
	return BaseWidgetInfoPropertiesLegendValuesEnum{
		LAST: BaseWidgetInfoPropertiesLegendValues{
			value: "last",
		},
		MAX: BaseWidgetInfoPropertiesLegendValues{
			value: "max",
		},
		MIN: BaseWidgetInfoPropertiesLegendValues{
			value: "min",
		},
		AVG: BaseWidgetInfoPropertiesLegendValues{
			value: "avg",
		},
		SUM: BaseWidgetInfoPropertiesLegendValues{
			value: "sum",
		},
	}
}

func (c BaseWidgetInfoPropertiesLegendValues) Value() string {
	return c.value
}

func (c BaseWidgetInfoPropertiesLegendValues) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoPropertiesLegendValues) UnmarshalJSON(b []byte) error {
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
