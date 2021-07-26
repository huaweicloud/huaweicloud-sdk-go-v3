package model

import (
	"encoding/json"

	"strings"
)

type ReclaimIndirectPartnerAccountReq struct {
	// 精英服务商ID。获取方法请参见查询精英服务商列表。

	IndirectPartnerId string `json:"indirect_partner_id"`
	// 回收金额。 华为云伙伴能力中心回收的精英服务商的账户金额。  说明： 回收金额不能大于精英服务商的账户余额。 单位：元。取值大于0且精确到小数点后2位。

	Amount float64 `json:"amount"`
}

func (o ReclaimIndirectPartnerAccountReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReclaimIndirectPartnerAccountReq struct{}"
	}

	return strings.Join([]string{"ReclaimIndirectPartnerAccountReq", string(data)}, " ")
}
