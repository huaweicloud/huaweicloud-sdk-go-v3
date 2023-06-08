package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type AdjustToIndirectPartnerReq struct {

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。
	IndirectPartnerId string `json:"indirect_partner_id"`

	// 华为云总经销商向云经销商拨款的金额。 单位：元。取值大于0且精确到小数点后2位。
	Amount *decimal.Decimal `json:"amount"`
}

func (o AdjustToIndirectPartnerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdjustToIndirectPartnerReq struct{}"
	}

	return strings.Join([]string{"AdjustToIndirectPartnerReq", string(data)}, " ")
}
