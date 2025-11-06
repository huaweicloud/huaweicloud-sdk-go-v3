package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BaseWidgetInfoRespProperties **参数解释** 额外信息
type BaseWidgetInfoRespProperties struct {

	// **参数解释** 聚合类型 **取值范围** 目前只有TopN这一种类型，折线图不支持该参数
	Filter *BaseWidgetInfoRespPropertiesFilter `json:"filter,omitempty"`

	// **参数解释** Top前N个；折线图时为随机展示的时序数据条数 **取值范围** 最小值为1，最大值为2147483647
	TopN *int32 `json:"topN,omitempty"`

	// **参数解释** 排序字段 **取值范围** - asc:正序 - desc:倒序
	Order *BaseWidgetInfoRespPropertiesOrder `json:"order,omitempty"`

	// **参数解释** 监控视图的描述信息 **取值范围** 长度为[0,200]个字符
	Description *string `json:"description,omitempty"`

	// **参数解释** 是否展示同比（上周同一时间）数据 **取值范围** - true:展示 - false:不展示
	LastWeekCompareEnable *bool `json:"last_week_compare_enable,omitempty"`

	// **参数解释** 是否展示环比（昨天同一时间）数据 **取值范围** - true:展示 - false:不展示
	YesterdayCompareEnable *bool `json:"yesterday_compare_enable,omitempty"`

	// **参数解释** 图例位置标记 **取值范围** - hide:表示隐藏图例 - right:表示图例放在监控视图右侧 - bottom:表示图例放在监控视图底部
	LegendLocation *BaseWidgetInfoRespPropertiesLegendLocation `json:"legend_location,omitempty"`

	// **参数解释** 当前时序数据需要在图例中展示的统计值名称列表
	LegendValues *[]BaseWidgetInfoRespPropertiesLegendValues `json:"legend_values,omitempty"`

	// **参数解释** 监控视图的阈值辅助线配置
	Thresholds *[]ThresholdInfoResp `json:"thresholds,omitempty"`

	// **参数解释** 同比环比总开关是否生效 **取值范围** - true:生效 - false:不生效
	IsAllCompareEnable *bool `json:"is_all_compare_enable,omitempty"`
}

func (o BaseWidgetInfoRespProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseWidgetInfoRespProperties struct{}"
	}

	return strings.Join([]string{"BaseWidgetInfoRespProperties", string(data)}, " ")
}

type BaseWidgetInfoRespPropertiesFilter struct {
	value string
}

type BaseWidgetInfoRespPropertiesFilterEnum struct {
	TOP_N BaseWidgetInfoRespPropertiesFilter
}

func GetBaseWidgetInfoRespPropertiesFilterEnum() BaseWidgetInfoRespPropertiesFilterEnum {
	return BaseWidgetInfoRespPropertiesFilterEnum{
		TOP_N: BaseWidgetInfoRespPropertiesFilter{
			value: "topN",
		},
	}
}

func (c BaseWidgetInfoRespPropertiesFilter) Value() string {
	return c.value
}

func (c BaseWidgetInfoRespPropertiesFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoRespPropertiesFilter) UnmarshalJSON(b []byte) error {
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

type BaseWidgetInfoRespPropertiesOrder struct {
	value string
}

type BaseWidgetInfoRespPropertiesOrderEnum struct {
	ASC  BaseWidgetInfoRespPropertiesOrder
	DESC BaseWidgetInfoRespPropertiesOrder
}

func GetBaseWidgetInfoRespPropertiesOrderEnum() BaseWidgetInfoRespPropertiesOrderEnum {
	return BaseWidgetInfoRespPropertiesOrderEnum{
		ASC: BaseWidgetInfoRespPropertiesOrder{
			value: "asc",
		},
		DESC: BaseWidgetInfoRespPropertiesOrder{
			value: "desc",
		},
	}
}

func (c BaseWidgetInfoRespPropertiesOrder) Value() string {
	return c.value
}

func (c BaseWidgetInfoRespPropertiesOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoRespPropertiesOrder) UnmarshalJSON(b []byte) error {
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

type BaseWidgetInfoRespPropertiesLegendLocation struct {
	value string
}

type BaseWidgetInfoRespPropertiesLegendLocationEnum struct {
	HIDE   BaseWidgetInfoRespPropertiesLegendLocation
	RIGHT  BaseWidgetInfoRespPropertiesLegendLocation
	BOTTOM BaseWidgetInfoRespPropertiesLegendLocation
}

func GetBaseWidgetInfoRespPropertiesLegendLocationEnum() BaseWidgetInfoRespPropertiesLegendLocationEnum {
	return BaseWidgetInfoRespPropertiesLegendLocationEnum{
		HIDE: BaseWidgetInfoRespPropertiesLegendLocation{
			value: "hide",
		},
		RIGHT: BaseWidgetInfoRespPropertiesLegendLocation{
			value: "right",
		},
		BOTTOM: BaseWidgetInfoRespPropertiesLegendLocation{
			value: "bottom",
		},
	}
}

func (c BaseWidgetInfoRespPropertiesLegendLocation) Value() string {
	return c.value
}

func (c BaseWidgetInfoRespPropertiesLegendLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoRespPropertiesLegendLocation) UnmarshalJSON(b []byte) error {
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

type BaseWidgetInfoRespPropertiesLegendValues struct {
	value string
}

type BaseWidgetInfoRespPropertiesLegendValuesEnum struct {
	LAST BaseWidgetInfoRespPropertiesLegendValues
	MAX  BaseWidgetInfoRespPropertiesLegendValues
	MIN  BaseWidgetInfoRespPropertiesLegendValues
	AVG  BaseWidgetInfoRespPropertiesLegendValues
	SUM  BaseWidgetInfoRespPropertiesLegendValues
}

func GetBaseWidgetInfoRespPropertiesLegendValuesEnum() BaseWidgetInfoRespPropertiesLegendValuesEnum {
	return BaseWidgetInfoRespPropertiesLegendValuesEnum{
		LAST: BaseWidgetInfoRespPropertiesLegendValues{
			value: "last",
		},
		MAX: BaseWidgetInfoRespPropertiesLegendValues{
			value: "max",
		},
		MIN: BaseWidgetInfoRespPropertiesLegendValues{
			value: "min",
		},
		AVG: BaseWidgetInfoRespPropertiesLegendValues{
			value: "avg",
		},
		SUM: BaseWidgetInfoRespPropertiesLegendValues{
			value: "sum",
		},
	}
}

func (c BaseWidgetInfoRespPropertiesLegendValues) Value() string {
	return c.value
}

func (c BaseWidgetInfoRespPropertiesLegendValues) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoRespPropertiesLegendValues) UnmarshalJSON(b []byte) error {
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
