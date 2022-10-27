package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExternalContactBase struct {

	// 其他号码。 > * 其他号码必须以国家码作为前缀 > * otherNumber填写时，otherNumberCountry也必须填写 > * 如果要清空手机号配置，则otherNumberCountry和otherNumber都要置为\"\"
	OtherNumber *string `json:"otherNumber,omitempty"`

	// [[其他号码所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	OtherNumberCountry *string `json:"otherNumberCountry,omitempty"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`

	// 手机号。 > * 手机号必须以国家码作为前缀 > * phone填写时，country也必须填写 > * 如果要清空手机号配置，则country和phone都要置为\"\"
	Phone *string `json:"phone,omitempty"`

	// 邮箱。
	Email *string `json:"email,omitempty"`

	// 公司名称。
	CorpName *string `json:"corpName,omitempty"`

	// 部门。
	DeptName *string `json:"deptName,omitempty"`

	// 职务。
	Position *string `json:"position,omitempty"`

	// 个人地址。
	Address *string `json:"address,omitempty"`

	// 备注。
	Remarks *string `json:"remarks,omitempty"`
}

func (o ExternalContactBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalContactBase struct{}"
	}

	return strings.Join([]string{"ExternalContactBase", string(data)}, " ")
}
