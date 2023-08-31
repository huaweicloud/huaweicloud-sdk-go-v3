package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowMonthUsageReq struct {

	// sim卡id列表，最多支持传入500个SIM卡id。sim_card_ids与iccids不能同时为空，sim_card_ids参数为空则根据iccids参数处理
	SimCardIds *[]int64 `json:"sim_card_ids,omitempty"`

	// iccid列表，最多支持传入500个iccid。sim_card_ids与iccids不能同时为空，sim_card_ids参数为空则根据iccids参数处理
	Iccids *[]string `json:"iccids,omitempty"`

	// 账期，最多支持传入本月在内的6个月账期，例如[2022-07, 2022-06]，不支持传入未来账期。
	BillingCycles []string `json:"billing_cycles"`
}

func (o ShowMonthUsageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonthUsageReq struct{}"
	}

	return strings.Join([]string{"ShowMonthUsageReq", string(data)}, " ")
}
