package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 联系方式, 应安全要求，放至body参数
type VerificationCodeDto struct {

	// 后台自动识别是手机号还是邮箱。 如果为手机号，必须加上国家码，例如中国大陆手机为“+86xxxxxxxxxxx”，当填写手机号时 “country”参数必填。 maxLength：255 minLength：1
	Contact string `json:"contact"`

	// 验证码，在校验的场景时需要携带
	VerificationCode *string `json:"verificationCode,omitempty"`

	// contact为手机号，则需带上手机号所属的国家。 例如国家为中国大陆则country参数取值为chinaPR 国家和国家码的对应关系请参考：https://support.huaweicloud.com/api-meeting/meeting_21_0109.html
	Country *string `json:"country,omitempty"`
}

func (o VerificationCodeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerificationCodeDto struct{}"
	}

	return strings.Join([]string{"VerificationCodeDto", string(data)}, " ")
}
