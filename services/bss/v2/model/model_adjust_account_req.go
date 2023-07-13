package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type AdjustAccountReq struct {

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。
	CustomerId string `json:"customer_id"`

	// 拨款金额。 单位：元。取值大于0且精确到小数点后2位。 注意该值不能大于“查询伙伴账户余额”接口响应消息表2中参数amount - designated_amount的值。
	Amount *decimal.Decimal `json:"amount"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。 云经销商（二级经销商）给子客户拨款时，需携带该参数。除此之外，该参数不做处理。  说明： 该参数存在的情况下，如果结果返回余额不足，表示对应的二级经销商的余额不足。如果该参数不存在，则余额不足表示调用的伙伴自身账号的余额不足。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o AdjustAccountReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdjustAccountReq struct{}"
	}

	return strings.Join([]string{"AdjustAccountReq", string(data)}, " ")
}
