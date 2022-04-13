package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerPostalAddressV2 struct {
	// 邮寄地址ID。

	AddressId *string `json:"address_id,omitempty"`
	// 收件人姓名。

	Recipient *string `json:"recipient,omitempty"`
	// 国家。例如：中国

	Nationality *string `json:"nationality,omitempty"`
	// 省/自治区/直辖市。例如：江苏省。

	Province *string `json:"province,omitempty"`
	// 市/区。例如：南京市。

	City *string `json:"city,omitempty"`
	// 区。例如：雨花区。

	District *string `json:"district,omitempty"`
	// 邮寄详细地址。

	Address *string `json:"address,omitempty"`
	// 邮编。

	Zipcode *string `json:"zipcode,omitempty"`
	// 国家码。例如： 中国：0086

	Areacode *string `json:"areacode,omitempty"`
	// 手机号码，不带国家码。

	MobilePhone *string `json:"mobile_phone,omitempty"`
	// 是否默认地址，默认为0。 1：默认地址0：非默认地址

	IsDefault *int32 `json:"is_default,omitempty"`
}

func (o CustomerPostalAddressV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerPostalAddressV2 struct{}"
	}

	return strings.Join([]string{"CustomerPostalAddressV2", string(data)}, " ")
}
