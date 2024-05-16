package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ViewBase 视图基础模型类。
type ViewBase struct {

	// 采集器名称。
	CollectorName *string `json:"collector_name,omitempty"`

	// 视图对应的指标集的名称。
	MetricSet *string `json:"metric_set,omitempty"`

	// 图标所需展示的标题。
	Title *string `json:"title,omitempty"`

	// 表格的方向，H：默认，表头横向，V：表头纵向。
	TableDirection *ViewBaseTableDirection `json:"table_direction,omitempty"`

	// 分组规则。
	GroupBy *string `json:"group_by,omitempty"`

	// 过滤列表模型。
	Filter *string `json:"filter,omitempty"`

	// 所需展示的字段列表模型。
	FieldItemList *[]FieldItem `json:"field_item_list,omitempty"`

	// span信息。
	Span *bool `json:"span,omitempty"`

	// span字段属性。
	SpanField *string `json:"span_field,omitempty"`

	// 排序规则。
	OrderBy *string `json:"order_by,omitempty"`

	// 是否只展示最近一笔数据。
	Latest *bool `json:"latest,omitempty"`

	// 视图类型。
	ViewType *ViewBaseViewType `json:"view_type,omitempty"`
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
