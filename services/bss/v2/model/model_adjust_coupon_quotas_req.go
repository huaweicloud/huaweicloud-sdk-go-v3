package model

import (
	"encoding/json"

	"strings"
)

type AdjustCouponQuotasReq struct {
	// 华为云伙伴能力中心发放的代金券额度的ID。

	QuotaId string `json:"quota_id"`
	// 精英服务商ID列表。

	IndirectPartnerIds []string `json:"indirect_partner_ids"`
	// 华为云伙伴能力中心向精英服务商发放的代金券额度值。 单位：元。取值大于0且精确到小数点后2位。

	QuotaAmount float64 `json:"quota_amount"`
}

func (o AdjustCouponQuotasReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AdjustCouponQuotasReq struct{}"
	}

	return strings.Join([]string{"AdjustCouponQuotasReq", string(data)}, " ")
}
