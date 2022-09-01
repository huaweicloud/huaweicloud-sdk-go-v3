package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 约束 - 号码和邮箱必须填一个，若企业未开启短信功能，则邮箱必填
type ActiveDto struct {

	// 手机号，如果为手机号，必须加上国家码。 例如中国大陆手机+86xxxxxxxxxxx，当填写手机号时 “country”参数必填。 maxLength：32 minLength：0
	SmsNumber *string `json:"smsNumber,omitempty" xml:"smsNumber"`

	// 若smsNumber为手机号,则需带上手机号所属的国家。 例如国家为中国大陆则country参数取值为chinaPR 国家和国家码的对应关系请参考：https://support.huaweicloud.com/api-meeting/meeting_21_0109.html
	Country *string `json:"country,omitempty" xml:"country"`

	// 邮件地址。 maxLength：255 minLength：0
	EmailAddr *string `json:"emailAddr,omitempty" xml:"emailAddr"`
}

func (o ActiveDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActiveDto struct{}"
	}

	return strings.Join([]string{"ActiveDto", string(data)}, " ")
}
