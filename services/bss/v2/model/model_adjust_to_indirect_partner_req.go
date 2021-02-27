package model

import (
	"encoding/json"

	"strings"
)

type AdjustToIndirectPartnerReq struct {
	// 精英服务商ID。您可以调用查询精英服务商列表接口获取。
	IndirectPartnerId string `json:"indirect_partner_id"`
	// 华为云伙伴能力中心向精英服务商拨款的金额。 单位：元。取值大于0且精确到小数点后2位。
	Amount float64 `json:"amount"`
}

func (o AdjustToIndirectPartnerReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AdjustToIndirectPartnerReq struct{}"
	}

	return strings.Join([]string{"AdjustToIndirectPartnerReq", string(data)}, " ")
}
