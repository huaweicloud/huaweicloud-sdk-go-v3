package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetPartnerDiscountsReq struct {
	// 精英服务商（二级经销商）ID。 精英服务商（二级经销商）给子客户设置折扣时需要携带这个字段。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
	// 客户折扣信息列表，最大支持10个。 具体请参见表1。

	SubCustomerDiscounts []SetSubCustomerDiscountV2 `json:"sub_customer_discounts"`
}

func (o SetPartnerDiscountsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPartnerDiscountsReq struct{}"
	}

	return strings.Join([]string{"SetPartnerDiscountsReq", string(data)}, " ")
}
