package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressItemDto struct {

	// 包含要显示的地址的格式化版本的字符串
	Formatted *string `json:"formatted,omitempty"`

	// 街道
	StreetAddress *string `json:"streetAddress,omitempty"`

	// 地址位置
	Locality *string `json:"locality,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 邮政编码
	PostalCode *string `json:"postalCode,omitempty"`

	// 国家或地区
	Country *string `json:"country,omitempty"`

	// 表示地址类型的字符串
	Type *string `json:"type,omitempty"`

	// 一个布尔值，表示这是否是用户的主地址
	Primary *bool `json:"primary,omitempty"`
}

func (o AddressItemDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressItemDto struct{}"
	}

	return strings.Join([]string{"AddressItemDto", string(data)}, " ")
}
