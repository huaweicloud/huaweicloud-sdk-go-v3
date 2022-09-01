package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 汇聚表格视图
type SumTableView struct {

	// 视图类型
	ViewType *SumTableViewViewType `json:"view_type,omitempty" xml:"view_type"`

	// 采集器名称
	CollectorName *string `json:"collector_name,omitempty" xml:"collector_name"`

	// 视图对应的指标集的名称
	MetricSet *string `json:"metric_set,omitempty" xml:"metric_set"`

	// 图标所需展示的标题
	Title *string `json:"title,omitempty" xml:"title"`

	// 表格的方向，H：默认，表头横向，V：表头纵向
	TableDirection *SumTableViewTableDirection `json:"table_direction,omitempty" xml:"table_direction"`

	// 分组规则
	GroupBy *string `json:"group_by,omitempty" xml:"group_by"`

	// 过滤列表模型
	Filter *string `json:"filter,omitempty" xml:"filter"`

	// 所需展示的字段列表模型列表
	FieldItemList *[]FieldItem `json:"field_item_list,omitempty" xml:"field_item_list"`

	// 跨度
	Span *bool `json:"span,omitempty" xml:"span"`

	// 跨度字段
	SpanField *string `json:"span_field,omitempty" xml:"span_field"`

	// 排序规则
	OrderBy *string `json:"order_by,omitempty" xml:"order_by"`

	// 是否只展示最近一笔数据
	Latest *bool `json:"latest,omitempty" xml:"latest"`
}

func (o SumTableView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SumTableView struct{}"
	}

	return strings.Join([]string{"SumTableView", string(data)}, " ")
}

type SumTableViewViewType struct {
	value string
}

type SumTableViewViewTypeEnum struct {
	TREND    SumTableViewViewType
	SUMTABLE SumTableViewViewType
	RAWTABLE SumTableViewViewType
}

func GetSumTableViewViewTypeEnum() SumTableViewViewTypeEnum {
	return SumTableViewViewTypeEnum{
		TREND: SumTableViewViewType{
			value: "trend",
		},
		SUMTABLE: SumTableViewViewType{
			value: "sumtable",
		},
		RAWTABLE: SumTableViewViewType{
			value: "rawtable",
		},
	}
}

func (c SumTableViewViewType) Value() string {
	return c.value
}

func (c SumTableViewViewType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SumTableViewViewType) UnmarshalJSON(b []byte) error {
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

type SumTableViewTableDirection struct {
	value string
}

type SumTableViewTableDirectionEnum struct {
	H SumTableViewTableDirection
	V SumTableViewTableDirection
}

func GetSumTableViewTableDirectionEnum() SumTableViewTableDirectionEnum {
	return SumTableViewTableDirectionEnum{
		H: SumTableViewTableDirection{
			value: "H",
		},
		V: SumTableViewTableDirection{
			value: "V",
		},
	}
}

func (c SumTableViewTableDirection) Value() string {
	return c.value
}

func (c SumTableViewTableDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SumTableViewTableDirection) UnmarshalJSON(b []byte) error {
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
