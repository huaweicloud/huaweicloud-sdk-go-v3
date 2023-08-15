package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActiveDto 设备联系人信息。
type ActiveDto struct {

	// 手机号。 例如中国大陆手机+86xxxxxxxxxxx。当填写手机号时 “country”参数必填。 > 号码和邮箱必须填一个，若企业未开启短信功能，则邮箱必填。
	SmsNumber *string `json:"smsNumber,omitempty"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`

	// 邮件地址。 > 号码和邮箱必须填一个，若企业未开启短信功能，则邮箱必填。
	EmailAddr *string `json:"emailAddr,omitempty"`
}

func (o ActiveDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActiveDto struct{}"
	}

	return strings.Join([]string{"ActiveDto", string(data)}, " ")
}
