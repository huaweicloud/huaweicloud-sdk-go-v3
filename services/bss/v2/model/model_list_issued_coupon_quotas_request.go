package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListIssuedCouponQuotasRequest struct {
	// 精英服务商的代金券额度ID。

	QuotaId *string `json:"quota_id,omitempty"`
	// 精英服务商ID。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
	// 父额度ID，即华为云伙伴能力中心用于发放给精英服务商代金券额度的额度ID。

	ParentQuotaId *string `json:"parent_quota_id,omitempty"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 每次查询记录数。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListIssuedCouponQuotasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListIssuedCouponQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListIssuedCouponQuotasRequest", string(data)}, " ")
}
