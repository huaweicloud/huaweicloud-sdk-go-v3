package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModAdminDto 待修改的管理员信息。
type ModAdminDto struct {

	// 名称。
	Name *string `json:"name,omitempty"`

	// 邮箱地址。
	Email *string `json:"email,omitempty"`

	// 手机号，必须加上国家码，例如中国大陆手机+86xxxxxxx，当填写手机号时 “country”参数必填。
	Phone *string `json:"phone,omitempty"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`
}

func (o ModAdminDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModAdminDto struct{}"
	}

	return strings.Join([]string{"ModAdminDto", string(data)}, " ")
}
