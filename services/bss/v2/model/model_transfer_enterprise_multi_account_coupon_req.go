package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TransferEnterpriseMultiAccountCouponReq struct {

	// 企业子账号的客户ID。您可以调用查询企业子账号列表接口，获取响应参数“id”的返回值。
	CustomerId string `json:"customer_id"`

	// 优惠券ID。您可以调用查询企业主账号可拨款优惠券列表接口，获取响应参数“coupon_id”的返回值。
	CouponId string `json:"coupon_id"`

	// 总划拨金额。单位为元。 单位：元。取值大于0且精确到小数点后2位。
	Amount string `json:"amount"`

	// 交易序列号，用于防止重复提交。 如果接口调用方不传此参数的值，则系统自动生成。如果接口调用方传入此参数的值，请采用UUID保证全局唯一。 此参数不携带或携带值为null或携带值为空串时，由系统自动生成。
	TransId *string `json:"trans_id,omitempty"`
}

func (o TransferEnterpriseMultiAccountCouponReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferEnterpriseMultiAccountCouponReq struct{}"
	}

	return strings.Join([]string{"TransferEnterpriseMultiAccountCouponReq", string(data)}, " ")
}
