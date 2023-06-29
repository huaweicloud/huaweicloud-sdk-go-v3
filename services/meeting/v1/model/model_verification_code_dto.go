package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerificationCodeDto 验证码信息。
type VerificationCodeDto struct {

	// 后台自动识别是手机号还是邮箱。 如果为手机号，必须加上国家码，例如中国大陆手机为“+86xxxxxxxxxxx”，当填写手机号时 “country”参数必填。
	Contact string `json:"contact"`

	// 验证码，在校验的场景时需要携带。
	VerificationCode *string `json:"verificationCode,omitempty"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`
}

func (o VerificationCodeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerificationCodeDto struct{}"
	}

	return strings.Join([]string{"VerificationCodeDto", string(data)}, " ")
}
