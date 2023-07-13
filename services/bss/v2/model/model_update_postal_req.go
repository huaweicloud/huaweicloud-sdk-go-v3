package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePostalReq struct {

	// 地址ID，地址的唯一索引。
	AddressId string `json:"address_id"`

	// 收件人姓名。 此参数不携带或携带值为null时，取值不更新。
	Recipient *string `json:"recipient,omitempty"`

	// 省/自治区/直辖市。例如：江苏省。 该参数携带并赋值时，city、district参数也必须赋值。 此参数不携带或携带值为null时，取值不更新。
	Province *string `json:"province,omitempty"`

	// 市/区。例如：南京市。 该参数携带并赋值时，province、district参数也必须赋值。 此参数不携带或携带值为null时，取值不更新。
	City *string `json:"city,omitempty"`

	// 区。例如：雨花区。 该参数携带并赋值时，province、city参数也必须赋值。 此参数不携带或携带值为null时，取值不更新。
	District *string `json:"district,omitempty"`

	// 邮寄详细地址。 此参数不携带或携带值为null时，取值不更新。
	Address *string `json:"address,omitempty"`

	// 邮寄地址所在邮编。 此参数不携带或携带值为null时，取值不更新。
	Zipcode *string `json:"zipcode,omitempty"`

	// 手机号码，不带国家码。 此参数不携带或携带值为null时，取值不更新。
	MobilePhone *string `json:"mobile_phone,omitempty"`

	// 是否默认地址，默认值为“0：非默认地址”。 1：默认地址0：非默认地址
	IsDefault *int32 `json:"is_default,omitempty"`
}

func (o UpdatePostalReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePostalReq struct{}"
	}

	return strings.Join([]string{"UpdatePostalReq", string(data)}, " ")
}
