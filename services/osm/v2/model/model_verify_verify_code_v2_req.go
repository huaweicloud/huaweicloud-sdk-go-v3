package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VerifyVerifyCodeV2Req struct {

	// 验证码
	VerifyCode string `json:"verify_code" xml:"verify_code"`

	// 联系方式的值
	ContactValue string `json:"contact_value" xml:"contact_value"`

	// 联系方式的类型。0：短信；1：邮件
	ContactWay int32 `json:"contact_way" xml:"contact_way"`

	// 国家码
	AreaCode *string `json:"area_code,omitempty" xml:"area_code"`

	// 子用户id
	XCustomerId *string `json:"x_customer_id,omitempty" xml:"x_customer_id"`

	// 客户id
	CustomerId *string `json:"customer_id,omitempty" xml:"customer_id"`
}

func (o VerifyVerifyCodeV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyVerifyCodeV2Req struct{}"
	}

	return strings.Join([]string{"VerifyVerifyCodeV2Req", string(data)}, " ")
}
