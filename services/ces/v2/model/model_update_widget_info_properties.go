package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateWidgetInfoProperties 视图展示配置
type UpdateWidgetInfoProperties struct {

	// 聚合类型，目前只有TopN这一种类型，折线图不支持该参数
	Filter *UpdateWidgetInfoPropertiesFilter `json:"filter,omitempty"`

	// Top值前N个;折线图时表示随机展示的时序数据条数
	TopN *int32 `json:"topN,omitempty"`

	// 排序字段，asc正序，desc倒序，折线图不支持该参数
	Order *UpdateWidgetInfoPropertiesOrder `json:"order,omitempty"`

	// 监控视图的描述信息
	Description *string `json:"description,omitempty"`

	// 是否展示同比（上周同一时间）数据，true:展示，false:不展示
	LastWeekCompareEnable *bool `json:"last_week_compare_enable,omitempty"`

	// 是否展示环比（昨天同一时间）数据，true:展示，false:不展示
	YesterdayCompareEnable *bool `json:"yesterday_compare_enable,omitempty"`

	// 图例位置标记，hide表示隐藏图例，right表示图例放在监控视图右侧，bottom表示图例放在监控视图底部，表格不支持该参数
	LegendLocation *UpdateWidgetInfoPropertiesLegendLocation `json:"legend_location,omitempty"`

	// 当前时序数据需要在图例中展示的统计值名称列表，表格不支持该参数，条形图和柱状图仅支持选择当前值
	LegendValues *[]UpdateWidgetInfoPropertiesLegendValues `json:"legend_values,omitempty"`

	// 监控视图的阈值辅助线配置
	Thresholds *[]ThresholdInfo `json:"thresholds,omitempty"`

	// 同比环比总开关是否生效;true:生效；false:不生效
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
