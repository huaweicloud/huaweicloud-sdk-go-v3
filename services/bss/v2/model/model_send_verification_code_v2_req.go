package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SendVerificationCodeV2Req struct {
	// 发送验证码的类型： 1：发送短信验证码

	ReceiverType int32 `json:"receiver_type"`
	// 发送验证码的超时时间。 如果不填的话，采用系统默认超时时间5分钟。 单位：分钟

	Timeout *int32 `json:"timeout,omitempty"`
	// 指定发送验证码的手机号。 目前系统只支持中国手机号。 示例：13XXXXXXXXX

	MobilePhone *string `json:"mobile_phone,omitempty"`
	// 根据该参数的取值选择发送短信验证码的语言。 zh-cn：中文en-us：英文

	Lang *string `json:"lang,omitempty"`
	// 验证码使用的场景，目前支持如下场景： 29：注册场景18：个人银行卡实名认证场景 不填写默认为29。

	Scene *int32 `json:"scene,omitempty"`
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。 如果scene=18的时候必填。

	CustomerId *string `json:"customer_id,omitempty"`
}

func (o SendVerificationCodeV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendVerificationCodeV2Req struct{}"
	}

	return strings.Join([]string{"SendVerificationCodeV2Req", string(data)}, " ")
}
