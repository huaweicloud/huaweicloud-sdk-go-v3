package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 待修改的管理员信息
type ModAdminDto struct {
	// 名称

	Name *string `json:"name,omitempty"`
	// 邮箱

	Email *string `json:"email,omitempty"`
	// 手机号，必须加上国家码，例如中国大陆手机+86xxxxxxx，当填写手机号时 “country”参数必填

	Phone *string `json:"phone,omitempty"`
	// 若smsNumber为手机号,则需带上手机号所属的国家。 例如国家为中国大陆则country参数取值为chinaPR 国家和国家码的对应关系请参考：https://support.huaweicloud.com/api-meeting/meeting_21_0109.html

	Country *string `json:"country,omitempty"`
}

func (o ModAdminDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModAdminDto struct{}"
	}

	return strings.Join([]string{"ModAdminDto", string(data)}, " ")
}
