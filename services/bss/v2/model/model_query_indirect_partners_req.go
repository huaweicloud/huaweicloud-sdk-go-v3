package model

import (
	"encoding/json"

	"strings"
)

type QueryIndirectPartnersReq struct {
	// 精英服务商伙伴的账号名。
	AccountName *string `json:"account_name,omitempty"`
	// 精英服务商关联华为云伙伴能力中心的开始时间。 UTC时间（包括时区），比如2016-03-28T00:00:00Z
	AssociatedOnBegin *string `json:"associated_on_begin,omitempty"`
	// 精英服务商关联华为云伙伴能力中心的结束时间。 UTC时间（包括时区），比如2016-03-28T00:00:00Z
	AssociatedOnEnd *string `json:"associated_on_end,omitempty"`
	// 偏移量，从0开始。默认值为0。
	Offset *int32 `json:"offset,omitempty"`
	// 每次查询的数量限制。默认值为10。
	Limit *int32 `json:"limit,omitempty"`
	// 精英服务商ID。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o QueryIndirectPartnersReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryIndirectPartnersReq struct{}"
	}

	return strings.Join([]string{"QueryIndirectPartnersReq", string(data)}, " ")
}
