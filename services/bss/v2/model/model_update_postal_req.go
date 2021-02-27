package model

import (
	"encoding/json"

	"strings"
)

type UpdatePostalReq struct {
	// 地址ID，地址的唯一索引。
	AddressId string `json:"address_id"`
	// 收件人姓名。
	Recipient *string `json:"recipient,omitempty"`
	// 省/自治区/直辖市。例如：江苏省。
	Province *string `json:"province,omitempty"`
	// 市/区。例如：南京市。
	City *string `json:"city,omitempty"`
	// 区。例如：雨花区。
	District *string `json:"district,omitempty"`
	// 邮寄详细地址。
	Address *string `json:"address,omitempty"`
	// 邮寄地址所在邮编。
	Zipcode *string `json:"zipcode,omitempty"`
	// 手机号码，不带国家码。
	MobilePhone *string `json:"mobile_phone,omitempty"`
	// 是否默认地址，默认为0。 1：默认地址0：非默认地址
	IsDefault *int32 `json:"is_default,omitempty"`
}

func (o UpdatePostalReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePostalReq struct{}"
	}

	return strings.Join([]string{"UpdatePostalReq", string(data)}, " ")
}
