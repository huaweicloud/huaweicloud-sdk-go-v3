package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPartnerBalancesRequest struct {
	// 精英服务商ID。获取方法请参见查询精英服务商列表。华为云伙伴能力中心（一级经销商）查询精英服务商（二级经销商）余额时，需要携带该参数；否则只能查询自己的账户余额。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListPartnerBalancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartnerBalancesRequest struct{}"
	}

	return strings.Join([]string{"ListPartnerBalancesRequest", string(data)}, " ")
}
