package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddPostalReq struct {

	// 收件人姓名。
	Recipient string `json:"recipient" xml:"recipient"`

	// 省、自治区或直辖市。例如：江苏省。
	Province string `json:"province" xml:"province"`

	// 市/区。例如：南京市。
	City string `json:"city" xml:"city"`

	// 区。例如：雨花台区。
	District string `json:"district" xml:"district"`

	// 邮寄详细地址。
	Address string `json:"address" xml:"address"`

	// 邮寄地址所在邮编。 此参数不携带或携带值为null时，不赋值。
	Zipcode *string `json:"zipcode,omitempty" xml:"zipcode"`

	// 手机号码，不带国家码。
	MobilePhone string `json:"mobile_phone" xml:"mobile_phone"`

	// 是否默认地址，默认值为“0：非默认地址”。 1：默认地址0：非默认地址
	IsDefault *int32 `json:"is_default,omitempty" xml:"is_default"`
}

func (o AddPostalReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPostalReq struct{}"
	}

	return strings.Join([]string{"AddPostalReq", string(data)}, " ")
}
