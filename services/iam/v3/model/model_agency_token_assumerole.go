package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AgencyTokenAssumerole struct {
	// 委托方A的账号ID。“domain_id”与“domain_name”至少填写一个。

	DomainId *string `json:"domain_id,omitempty"`
	// 委托方A的账号名称。“domain_id”与“domain_name”至少填写一个。

	DomainName *string `json:"domain_name,omitempty"`
	// 委托方A创建的委托的名称。

	AgencyName string `json:"agency_name"`
}

func (o AgencyTokenAssumerole) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyTokenAssumerole struct{}"
	}

	return strings.Join([]string{"AgencyTokenAssumerole", string(data)}, " ")
}
