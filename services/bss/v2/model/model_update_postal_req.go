package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePostalReq struct {

	// 地址ID，地址的唯一索引。
	AddressId string `json:"address_id" xml:"address_id"`

	// 收件人姓名。 此参数不携带或携带值为null时，取值不更新。
	Recipient *string `json:"recipient,omitempty" xml:"recipient"`

	// 省/自治区/直辖市。例如：江苏省。 此参数不携带或携带值为null时，取值不更新。
	Province *string `json:"province,omitempty" xml:"province"`

	// 市/区。例如：南京市。 此参数不携带或携带值为null时，取值不更新。
	City *string `json:"city,omitempty" xml:"city"`

	// 区。例如：雨花区。 此参数不携带或携带值为null时，取值不更新。
	District *string `json:"district,omitempty" xml:"district"`

	// 邮寄详细地址。 此参数不携带或携带值为null时，取值不更新。
	Address *string `json:"address,omitempty" xml:"address"`

	// 邮寄地址所在邮编。 此参数不携带或携带值为null时，取值不更新。
	Zipcode *string `json:"zipcode,omitempty" xml:"zipcode"`

	// 手机号码，不带国家码。 此参数不携带或携带值为null时，取值不更新。
	MobilePhone *string `json:"mobile_phone,omitempty" xml:"mobile_phone"`

	// 是否默认地址，默认值为“0：非默认地址”。 1：默认地址0：非默认地址
	IsDefault *int32 `json:"is_default,omitempty" xml:"is_default"`
}

func (o UpdatePostalReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePostalReq struct{}"
	}

	return strings.Join([]string{"UpdatePostalReq", string(data)}, " ")
}
