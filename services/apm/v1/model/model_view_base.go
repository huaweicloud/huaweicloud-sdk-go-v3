package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 视图基础模型类
type ViewBase struct {

	// 采集器名称
	CollectorName *string `json:"collector_name,omitempty" xml:"collector_name"`

	// 视图对应的指标集的名称
	MetricSet *string `json:"metric_set,omitempty" xml:"metric_set"`

	// 图标所需展示的标题
	Title *string `json:"title,omitempty" xml:"title"`

	// 表格的方向，H：默认，表头横向，V：表头纵向
	TableDirection *ViewBaseTableDirection `json:"table_direction,omitempty" xml:"table_direction"`

	// 分组规则
	GroupBy *string `json:"group_by,omitempty" xml:"group_by"`

	// 过滤列表模型
	Filter *string `json:"filter,omitempty" xml:"filter"`

	// 所需展示的字段列表模型
	FieldItemList *[]FieldItem `json:"field_item_list,omitempty" xml:"field_item_list"`

	// span信息
	Span *bool `json:"span,omitempty" xml:"span"`

	// span字段属性
	SpanField *string `json:"span_field,omitempty" xml:"span_field"`

	// 排序规则
	OrderBy *string `json:"order_by,omitempty" xml:"order_by"`

	// 是否只展示最近一笔数据
	Latest *string `json:"latest,omitempty" xml:"latest"`

	// 视图类型
	ViewType *ViewBaseViewType `json:"view_type,omitempty" xml:"view_type"`
}

func (o ViewBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ViewBase struct{}"
	}

	return strings.Join([]string{"ViewBase", string(data)}, " ")
}

type ViewBaseTableDirection struct {
	value string
}

type ViewBaseTableDirectionEnum struct {
	H ViewBaseTableDirection
	V ViewBaseTableDirection
}

func GetViewBaseTableDirectionEnum() ViewBaseTableDirectionEnum {
	return ViewBaseTableDirectionEnum{
		H: ViewBaseTableDirection{
			value: "H",
		},
		V: ViewBaseTableDirection{
			value: "V",
		},
	}
}

func (c ViewBaseTableDirection) Value() string {
	return c.value
}

func (c ViewBaseTableDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ViewBaseTableDirection) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ViewBaseViewType struct {
	value string
}

type ViewBaseViewTypeEnum struct {
	TREND    ViewBaseViewType
	SUMTABLE ViewBaseViewType
	RAWTABLE ViewBaseViewType
}

func GetViewBaseViewTypeEnum() ViewBaseViewTypeEnum {
	return ViewBaseViewTypeEnum{
		TREND: ViewBaseViewType{
			value: "trend",
		},
		SUMTABLE: ViewBaseViewType{
			value: "sumtable",
		},
		RAWTABLE: ViewBaseViewType{
			value: "rawtable",
		},
	}
}

func (c ViewBaseViewType) Value() string {
	return c.value
}

func (c ViewBaseViewType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ViewBaseViewType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
