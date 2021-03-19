package model

import (
	"encoding/json"

	"strings"
)

type IndirectPartnerInfo struct {
	// 精英服务商ID。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
	// 精英服务商的手机号码。

	MobilePhone *string `json:"mobile_phone,omitempty"`
	// 精英服务商的邮箱。

	Email *string `json:"email,omitempty"`
	// 精英服务商的账户名。

	AccountName *string `json:"account_name,omitempty"`
	// 精英服务商的名称。

	Name *string `json:"name,omitempty"`
	// 精英服务商关联华为云伙伴能力中心的时间。 UTC时间（包括时区），比如2016-03-28T00:00:00Z

	AssociatedOn *string `json:"associated_on,omitempty"`
}

func (o IndirectPartnerInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IndirectPartnerInfo struct{}"
	}

	return strings.Join([]string{"IndirectPartnerInfo", string(data)}, " ")
}
