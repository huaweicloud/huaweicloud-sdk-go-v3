package model

import (
	"encoding/json"

	"strings"
)

type QueryCustomersBalancesReq struct {
	// 客户信息列表。 具体请参见表1。
	CustomerInfos []CustomerInfoV2 `json:"customer_infos"`
	// 精英服务商ID。 华为云伙伴能力中心（一级经销商）查询精英服务商（二级经销商）子客户的账户余额时，需要携带该参数。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o QueryCustomersBalancesReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryCustomersBalancesReq struct{}"
	}

	return strings.Join([]string{"QueryCustomersBalancesReq", string(data)}, " ")
}
