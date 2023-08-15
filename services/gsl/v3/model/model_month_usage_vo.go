package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MonthUsageVo struct {

	// SIM卡ID
	SimCardId *int64 `json:"sim_card_id,omitempty"`

	// 月用量
	FlowUsages *[]FlowUsageVo `json:"flow_usages,omitempty"`
}

func (o MonthUsageVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonthUsageVo struct{}"
	}

	return strings.Join([]string{"MonthUsageVo", string(data)}, " ")
}
