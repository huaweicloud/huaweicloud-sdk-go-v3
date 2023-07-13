package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReclaimCouponQuotasReq struct {

	// 被回收的云经销商的代金券额度ID。获取方法请参见查询优惠券额度。
	QuotaIds []string `json:"quota_ids"`

	// 回收时的备注。 此参数不携带或携带值为null时，不赋值；携带值为空串时，赋值空串。
	Remark *string `json:"remark,omitempty"`
}

func (o ReclaimCouponQuotasReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimCouponQuotasReq struct{}"
	}

	return strings.Join([]string{"ReclaimCouponQuotasReq", string(data)}, " ")
}
