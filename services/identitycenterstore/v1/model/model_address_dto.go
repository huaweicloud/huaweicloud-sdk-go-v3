package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddressDto The address associated with the specified user.
type AddressDto struct {

	// 国家/地区
	Country *string `json:"country,omitempty"`

	// 包含要显示的地址的格式化版本的字符串
	Formatted *string `json:"formatted,omitempty"`

	// 地址位置
	Locality *string `json:"locality,omitempty"`

	// 邮政编码
	PostalCode *string `json:"postal_code,omitempty"`

	// 一个布尔值，表示这是否为用户的主地址
	Primary *bool `json:"primary,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 街道
	StreetAddress *string `json:"street_address,omitempty"`

	// 表示地址类型的字符串
	Type *string `json:"type,omitempty"`
}

func (o AddressDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressDto struct{}"
	}

	return strings.Join([]string{"AddressDto", string(data)}, " ")
}
