package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 指标分子过滤条件
type MetricRequest3Dividend struct {

	// 是否按时
	OnTime *string `json:"on_time,omitempty" xml:"on_time"`

	// 自定义字段
	CustomField16 *string `json:"custom_field16,omitempty" xml:"custom_field16"`

	// 自定义字段
	CustomField17 *string `json:"custom_field17,omitempty" xml:"custom_field17"`

	// 自定义字段
	CustomField18 *string `json:"custom_field18,omitempty" xml:"custom_field18"`

	// 自定义字段
	CustomField19 *string `json:"custom_field19,omitempty" xml:"custom_field19"`

	// 自定义字段
	CustomField20 *string `json:"custom_field20,omitempty" xml:"custom_field20"`

	// 自定义字段
	CustomField21 *string `json:"custom_field21,omitempty" xml:"custom_field21"`

	// 自定义字段
	CustomField22 *string `json:"custom_field22,omitempty" xml:"custom_field22"`

	// 自定义字段
	CustomField23 *string `json:"custom_field23,omitempty" xml:"custom_field23"`

	// 自定义字段
	CustomField24 *string `json:"custom_field24,omitempty" xml:"custom_field24"`

	// 自定义字段
	CustomField25 *string `json:"custom_field25,omitempty" xml:"custom_field25"`

	// 自定义字段
	CustomField26 *string `json:"custom_field26,omitempty" xml:"custom_field26"`

	// 自定义字段
	CustomField27 *string `json:"custom_field27,omitempty" xml:"custom_field27"`

	// 自定义字段
	CustomField28 *string `json:"custom_field28,omitempty" xml:"custom_field28"`

	// 自定义字段
	CustomField29 *string `json:"custom_field29,omitempty" xml:"custom_field29"`

	// 自定义字段
	CustomField30 *string `json:"custom_field30,omitempty" xml:"custom_field30"`

	// 自定义字段
	CustomField31 *string `json:"custom_field31,omitempty" xml:"custom_field31"`

	// 自定义字段
	CustomField32 *string `json:"custom_field32,omitempty" xml:"custom_field32"`

	// 自定义字段
	CustomField33 *string `json:"custom_field33,omitempty" xml:"custom_field33"`

	// 自定义字段
	CustomField34 *string `json:"custom_field34,omitempty" xml:"custom_field34"`

	// 自定义字段
	CustomField35 *string `json:"custom_field35,omitempty" xml:"custom_field35"`

	// 自定义字段
	CustomField36 *string `json:"custom_field36,omitempty" xml:"custom_field36"`

	// 自定义字段
	CustomField37 *string `json:"custom_field37,omitempty" xml:"custom_field37"`

	// 自定义字段
	CustomField38 *string `json:"custom_field38,omitempty" xml:"custom_field38"`

	// 自定义字段
	CustomField39 *string `json:"custom_field39,omitempty" xml:"custom_field39"`

	// 自定义字段
	CustomField40 *string `json:"custom_field40,omitempty" xml:"custom_field40"`
}

func (o MetricRequest3Dividend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricRequest3Dividend struct{}"
	}

	return strings.Join([]string{"MetricRequest3Dividend", string(data)}, " ")
}
