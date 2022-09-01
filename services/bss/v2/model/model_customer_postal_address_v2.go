package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerPostalAddressV2 struct {

	// 邮寄地址ID。
	AddressId *string `json:"address_id,omitempty" xml:"address_id"`

	// 收件人姓名。
	Recipient *string `json:"recipient,omitempty" xml:"recipient"`

	// 国家。例如：中国
	Nationality *string `json:"nationality,omitempty" xml:"nationality"`

	// 省/自治区/直辖市。例如：江苏省。
	Province *string `json:"province,omitempty" xml:"province"`

	// 市/区。例如：南京市。
	City *string `json:"city,omitempty" xml:"city"`

	// 区。例如：雨花区。
	District *string `json:"district,omitempty" xml:"district"`

	// 邮寄详细地址。
	Address *string `json:"address,omitempty" xml:"address"`

	// 邮编。
	Zipcode *string `json:"zipcode,omitempty" xml:"zipcode"`

	// 国家码。例如： 中国：0086
	Areacode *string `json:"areacode,omitempty" xml:"areacode"`

	// 手机号码，不带国家码。
	MobilePhone *string `json:"mobile_phone,omitempty" xml:"mobile_phone"`

	// 是否默认地址，默认值为“0：非默认地址”。 1：默认地址0：非默认地址
	IsDefault *int32 `json:"is_default,omitempty" xml:"is_default"`
}

func (o CustomerPostalAddressV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerPostalAddressV2 struct{}"
	}

	return strings.Join([]string{"CustomerPostalAddressV2", string(data)}, " ")
}
