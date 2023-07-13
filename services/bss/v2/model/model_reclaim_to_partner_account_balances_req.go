package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type ReclaimToPartnerAccountBalancesReq struct {

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。
	CustomerId string `json:"customer_id"`

	// 回收的金额。 单位：元。取值大于0且精确到小数点后2位。
	Amount *decimal.Decimal `json:"amount"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。云经销商（二级经销商）回收云经销商（二级经销商）的子客户账户余额时，需携带此参数；除此之外，该参数不做处理；否则只能回收自己的子客户账户余额。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ReclaimToPartnerAccountBalancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimToPartnerAccountBalancesReq struct{}"
	}

	return strings.Join([]string{"ReclaimToPartnerAccountBalancesReq", string(data)}, " ")
}
