package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type AdjustCouponQuotasReq struct {

	// 华为云总经销商发放的代金券额度的ID。获取方法请参见查询优惠券额度。
	QuotaId string `json:"quota_id"`

	// 云经销商ID列表。
	IndirectPartnerIds []string `json:"indirect_partner_ids"`

	// 华为云总经销商向云经销商发放的代金券额度值。 单位：元。取值大于0且精确到小数点后2位。
	QuotaAmount *decimal.Decimal `json:"quota_amount"`
}

func (o AdjustCouponQuotasReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdjustCouponQuotasReq struct{}"
	}

	return strings.Join([]string{"AdjustCouponQuotasReq", string(data)}, " ")
}
