package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricRequest2Dividend 指标分子过滤条件
type MetricRequest2Dividend struct {

	// 自定义字段
	CustomField16 *string `json:"custom_field16,omitempty"`

	// 自定义字段
	CustomField17 *string `json:"custom_field17,omitempty"`

	// 自定义字段
	CustomField18 *string `json:"custom_field18,omitempty"`

	// 自定义字段
	CustomField19 *string `json:"custom_field19,omitempty"`

	// 自定义字段
	CustomField20 *string `json:"custom_field20,omitempty"`

	// 自定义字段
	CustomField21 *string `json:"custom_field21,omitempty"`

	// 自定义字段
	CustomField22 *string `json:"custom_field22,omitempty"`

	// 自定义字段
	CustomField23 *string `json:"custom_field23,omitempty"`

	// 自定义字段
	CustomField24 *string `json:"custom_field24,omitempty"`

	// 自定义字段
	CustomField25 *string `json:"custom_field25,omitempty"`

	// 自定义字段
	CustomField26 *string `json:"custom_field26,omitempty"`

	// 自定义字段
	CustomField27 *string `json:"custom_field27,omitempty"`

	// 自定义字段
	CustomField28 *string `json:"custom_field28,omitempty"`

	// 自定义字段
	CustomField29 *string `json:"custom_field29,omitempty"`

	// 自定义字段
	CustomField30 *string `json:"custom_field30,omitempty"`

	// 自定义字段
	CustomField31 *string `json:"custom_field31,omitempty"`

	// 自定义字段
	CustomField32 *string `json:"custom_field32,omitempty"`

	// 自定义字段
	CustomField33 *string `json:"custom_field33,omitempty"`

	// 自定义字段
	CustomField34 *string `json:"custom_field34,omitempty"`

	// 自定义字段
	CustomField35 *string `json:"custom_field35,omitempty"`

	// 自定义字段
	CustomField36 *string `json:"custom_field36,omitempty"`

	// 自定义字段
	CustomField37 *string `json:"custom_field37,omitempty"`

	// 自定义字段
	CustomField38 *string `json:"custom_field38,omitempty"`

	// 自定义字段
	CustomField39 *string `json:"custom_field39,omitempty"`

	// 自定义字段
	CustomField40 *string `json:"custom_field40,omitempty"`
}

func (o MetricRequest2Dividend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricRequest2Dividend struct{}"
	}

	return strings.Join([]string{"MetricRequest2Dividend", string(data)}, " ")
}
