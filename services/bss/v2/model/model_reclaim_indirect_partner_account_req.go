package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type ReclaimIndirectPartnerAccountReq struct {

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。
	IndirectPartnerId string `json:"indirect_partner_id"`

	// 回收金额。 华为云总经销商回收的云经销商的账户金额。  说明： 回收金额不能大于云经销商的账户余额。 单位：元。取值大于0且精确到小数点后2位。
	Amount *decimal.Decimal `json:"amount"`
}

func (o ReclaimIndirectPartnerAccountReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimIndirectPartnerAccountReq struct{}"
	}

	return strings.Join([]string{"ReclaimIndirectPartnerAccountReq", string(data)}, " ")
}
