package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NewCustomerV2 struct {

	// 企业子账号登录名。
	CustomerName string `json:"customer_name"`

	// 企业管理员的手机号码。如果use_pri_mobile_phone取值为Y，则这个参数无效，否则必选。
	MobilePhone *string `json:"mobile_phone,omitempty"`

	// 是否使用企业主账号手机号码作为子账号手机号码： Y：是N：否（默认值） 当为Y时，mobile_phone输入无效。
	UsePriMobilePhone *string `json:"use_pri_mobile_phone,omitempty"`

	// 企业子账号的登录密码。
	Password string `json:"password"`

	// 验证码，只有输入企业子账号的手机号的情况下，才需要填写该字段。 具体请参见发送短信验证码。
	VerificationCode *string `json:"verification_code,omitempty"`
}

func (o NewCustomerV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NewCustomerV2 struct{}"
	}

	return strings.Join([]string{"NewCustomerV2", string(data)}, " ")
}
