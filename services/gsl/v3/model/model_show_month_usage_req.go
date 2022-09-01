package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowMonthUsageReq struct {

	// sim卡id列表
	SimCardIds []int64 `json:"sim_card_ids" xml:"sim_card_ids"`

	// 账期
	BillingCycles []string `json:"billing_cycles" xml:"billing_cycles"`
}

func (o ShowMonthUsageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonthUsageReq struct{}"
	}

	return strings.Join([]string{"ShowMonthUsageReq", string(data)}, " ")
}
