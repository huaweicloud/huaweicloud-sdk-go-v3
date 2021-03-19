package model

import (
	"encoding/json"

	"strings"
)

type ReclaimCouponQuotasReq struct {
	// 被回收的精英服务商的代金券额度ID。

	QuotaIds []string `json:"quota_ids"`
	// 回收时的备注。

	Remark *string `json:"remark,omitempty"`
}

func (o ReclaimCouponQuotasReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReclaimCouponQuotasReq struct{}"
	}

	return strings.Join([]string{"ReclaimCouponQuotasReq", string(data)}, " ")
}
