package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 待修改的企业基本信息，不带或参数null，则不修改
type ModCorpBasicDto struct {
	// 企业名称，格式必须满足^[^#%&'+;<>=\"'？?\\\\……/]*$

	Name *string `json:"name,omitempty"`
	// 手机号，必须加上国家码，例如中国大陆手机+86xxxxxxx，当填写手机号时 “country”参数必填,手机格式必须满足(^$|^[+]?[0-9]+$)

	Phone *string `json:"phone,omitempty"`
	// 若smsNumber为手机号,则需带上手机号所属的国家。 例如国家为中国大陆则country参数取值为chinaPR 国家和国家码的对应关系请参考：https://support.huaweicloud.com/api-meeting/meeting_21_0109.html

	Country *string `json:"country,omitempty"`
	// 传真号码,格式必须满足(^$|^[+]?[0-9]+$)

	Fax *string `json:"fax,omitempty"`
	// 邮箱地址,格式必须满足(^$|^[\\w-+]+(\\.[\\w-+]+)*@[\\w-]+(\\.[\\w-]+)*(\\.[\\w-]{1,})$)

	Email *string `json:"email,omitempty"`
	// 地址

	Address *string `json:"address,omitempty"`
	// 备注

	Description *string `json:"description,omitempty"`
}

func (o ModCorpBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModCorpBasicDto struct{}"
	}

	return strings.Join([]string{"ModCorpBasicDto", string(data)}, " ")
}
