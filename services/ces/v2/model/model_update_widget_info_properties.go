package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateWidgetInfoProperties **参数解释** 视图展示配置 **约束限制** 不涉及
type UpdateWidgetInfoProperties struct {

	// **参数解释** 聚合类型 **约束限制** 折线图不支持该参数 **取值范围** 枚举值： - topN topN类型 **默认取值** 不涉及
	Filter *UpdateWidgetInfoPropertiesFilter `json:"filter,omitempty"`

	// **参数解释** Top值前N个 **约束限制** 折线图时表示随机展示的时序数据条数 **取值范围** Top值为[1,2147483647] **默认取值** 不涉及
	TopN *int32 `json:"topN,omitempty"`

	// **参数解释** 排序字段 **约束限制** 折线图不支持该参数 **取值范围** 枚举值： - asc 正序 - desc 倒序 **默认取值** 不涉及
	Order *UpdateWidgetInfoPropertiesOrder `json:"order,omitempty"`

	// **参数解释** 监控视图的描述信息 **约束限制** 不涉及 **取值范围** 信息长度为[0,200]个字符 **默认取值** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释** 是否展示同比（上周同一时间）数据 **约束限制** 不涉及 **取值范围** - true 展示 - false 不展示 **默认取值** 不涉及
	LastWeekCompareEnable *bool `json:"last_week_compare_enable,omitempty"`

	// **参数解释** 是否展示环比（昨天同一时间）数据 **约束限制** 不涉及 **取值范围** - true 展示 - false 不展示 **默认取值** 不涉及
	YesterdayCompareEnable *bool `json:"yesterday_compare_enable,omitempty"`

	// **参数解释** 图例位置标记 **约束限制** 表格不支持该参数 **取值范围** 枚举值： - hide 隐藏图例 - right 图例放在监控视图右侧 - bottom 图例放在监控视图底部 **默认取值** 不涉及
	LegendLocation *UpdateWidgetInfoPropertiesLegendLocation `json:"legend_location,omitempty"`

	// **参数解释** 当前时序数据需要在图例中展示的统计值名称列表 **约束限制** 表格不支持该参数，条形图和柱状图仅支持选择当前值 包含的名称对象个数为[0,5]
	LegendValues *[]UpdateWidgetInfoPropertiesLegendValues `json:"legend_values,omitempty"`

	// **参数解释**   监控视图的阈值辅助线配置   **约束限制**   包含的配置对象个数为[0,6]
	Thresholds *[]ThresholdInfo `json:"thresholds,omitempty"`

	// **参数解释** 同比环比总开关是否生效 **约束限制** 不涉及 **取值范围** - true 生效 - false 不生效 **默认取值** 不涉及
	IsAllCompareEnable *bool `json:"is_all_compare_enable,omitempty"`
}

func (o UpdateWidgetInfoProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWidgetInfoProperties struct{}"
	}

	return strings.Join([]string{"UpdateWidgetInfoProperties", string(data)}, " ")
}

type UpdateWidgetInfoPropertiesFilter struct {
	value string
}

type UpdateWidgetInfoPropertiesFilterEnum struct {
	TOP_N UpdateWidgetInfoPropertiesFilter
}

func GetUpdateWidgetInfoPropertiesFilterEnum() UpdateWidgetInfoPropertiesFilterEnum {
	return UpdateWidgetInfoPropertiesFilterEnum{
		TOP_N: UpdateWidgetInfoPropertiesFilter{
			value: "topN",
		},
	}
}

func (c UpdateWidgetInfoPropertiesFilter) Value() string {
	return c.value
}

func (c UpdateWidgetInfoPropertiesFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateWidgetInfoPropertiesFilter) UnmarshalJSON(b []byte) error {
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

type UpdateWidgetInfoPropertiesOrder struct {
	value string
}

type UpdateWidgetInfoPropertiesOrderEnum struct {
	ASC  UpdateWidgetInfoPropertiesOrder
	DESC UpdateWidgetInfoPropertiesOrder
}

func GetUpdateWidgetInfoPropertiesOrderEnum() UpdateWidgetInfoPropertiesOrderEnum {
	return UpdateWidgetInfoPropertiesOrderEnum{
		ASC: UpdateWidgetInfoPropertiesOrder{
			value: "asc",
		},
		DESC: UpdateWidgetInfoPropertiesOrder{
			value: "desc",
		},
	}
}

func (c UpdateWidgetInfoPropertiesOrder) Value() string {
	return c.value
}

func (c UpdateWidgetInfoPropertiesOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateWidgetInfoPropertiesOrder) UnmarshalJSON(b []byte) error {
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

type UpdateWidgetInfoPropertiesLegendLocation struct {
	value string
}

type UpdateWidgetInfoPropertiesLegendLocationEnum struct {
	HIDE   UpdateWidgetInfoPropertiesLegendLocation
	RIGHT  UpdateWidgetInfoPropertiesLegendLocation
	BOTTOM UpdateWidgetInfoPropertiesLegendLocation
}

func GetUpdateWidgetInfoPropertiesLegendLocationEnum() UpdateWidgetInfoPropertiesLegendLocationEnum {
	return UpdateWidgetInfoPropertiesLegendLocationEnum{
		HIDE: UpdateWidgetInfoPropertiesLegendLocation{
			value: "hide",
		},
		RIGHT: UpdateWidgetInfoPropertiesLegendLocation{
			value: "right",
		},
		BOTTOM: UpdateWidgetInfoPropertiesLegendLocation{
			value: "bottom",
		},
	}
}

func (c UpdateWidgetInfoPropertiesLegendLocation) Value() string {
	return c.value
}

func (c UpdateWidgetInfoPropertiesLegendLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateWidgetInfoPropertiesLegendLocation) UnmarshalJSON(b []byte) error {
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

type UpdateWidgetInfoPropertiesLegendValues struct {
	value string
}

type UpdateWidgetInfoPropertiesLegendValuesEnum struct {
	LAST UpdateWidgetInfoPropertiesLegendValues
	MAX  UpdateWidgetInfoPropertiesLegendValues
	MIN  UpdateWidgetInfoPropertiesLegendValues
	AVG  UpdateWidgetInfoPropertiesLegendValues
	SUM  UpdateWidgetInfoPropertiesLegendValues
}

func GetUpdateWidgetInfoPropertiesLegendValuesEnum() UpdateWidgetInfoPropertiesLegendValuesEnum {
	return UpdateWidgetInfoPropertiesLegendValuesEnum{
		LAST: UpdateWidgetInfoPropertiesLegendValues{
			value: "last",
		},
		MAX: UpdateWidgetInfoPropertiesLegendValues{
			value: "max",
		},
		MIN: UpdateWidgetInfoPropertiesLegendValues{
			value: "min",
		},
		AVG: UpdateWidgetInfoPropertiesLegendValues{
			value: "avg",
		},
		SUM: UpdateWidgetInfoPropertiesLegendValues{
			value: "sum",
		},
	}
}

func (c UpdateWidgetInfoPropertiesLegendValues) Value() string {
	return c.value
}

func (c UpdateWidgetInfoPropertiesLegendValues) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateWidgetInfoPropertiesLegendValues) UnmarshalJSON(b []byte) error {
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
