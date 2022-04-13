package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReclaimToPartnerAccountBalancesReq struct {
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。

	CustomerId string `json:"customer_id"`
	// 回收的金额。 单位：元。取值大于0且精确到小数点后2位。

	Amount float64 `json:"amount"`
	// 精英服务商ID。获取方法请参见查询精英服务商列表。 华为云伙伴能力中心（一级经销商）回收精英服务商（二级经销商）的子客户账户余额时，需携带此参数；否则只能回收自己的子客户账户余额。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ReclaimToPartnerAccountBalancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimToPartnerAccountBalancesReq struct{}"
	}

	return strings.Join([]string{"ReclaimToPartnerAccountBalancesReq", string(data)}, " ")
}
