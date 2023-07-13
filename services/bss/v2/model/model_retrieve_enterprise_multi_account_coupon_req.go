package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetrieveEnterpriseMultiAccountCouponReq struct {

	// 企业子账号的客户ID。您可以调用查询企业子账号列表接口，获取响应参数“id”的返回值。
	CustomerId string `json:"customer_id"`

	// 优惠券ID。您可以调用查询企业子账号可回收优惠券列表接口，获取响应参数“coupon_id”的返回值。
	CouponId string `json:"coupon_id"`

	// 子优惠券ID。您可以调用查询企业子账号可回收优惠券列表接口，获取响应参数“sub_coupon_id”的返回值。
	SubCouponId string `json:"sub_coupon_id"`

	// 交易序列号，用于防止重复提交。 如果接口调用方不传，则企业管理微服务后台生成如果接口调用方传入，则请采用UUID保证全局唯一 此参数不携带或携带值为null或携带值为空串时，由系统自动生成。
	TransId *string `json:"trans_id,omitempty"`
}

func (o RetrieveEnterpriseMultiAccountCouponReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetrieveEnterpriseMultiAccountCouponReq struct{}"
	}

	return strings.Join([]string{"RetrieveEnterpriseMultiAccountCouponReq", string(data)}, " ")
}
