package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IndirectPartnerInfo struct {

	// 云经销商ID。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`

	// 云经销商的手机号码。
	MobilePhone *string `json:"mobile_phone,omitempty"`

	// 云经销商的邮箱。
	Email *string `json:"email,omitempty"`

	// 云经销商的账户名。
	AccountName *string `json:"account_name,omitempty"`

	// 云经销商的名称。
	Name *string `json:"name,omitempty"`

	// 云经销商关联华为云总经销商的时间。 UTC时间（包括时区），例如2016-03-28T00:00:00Z。
	AssociatedOn *string `json:"associated_on,omitempty"`

	// 客户经理登录账户名。
	AccountManagerId *string `json:"account_manager_id,omitempty"`

	// 客户经理的名称。
	AccountManagerName *string `json:"account_manager_name,omitempty"`
}

func (o IndirectPartnerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndirectPartnerInfo struct{}"
	}

	return strings.Join([]string{"IndirectPartnerInfo", string(data)}, " ")
}
