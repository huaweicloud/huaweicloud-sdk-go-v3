package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimCardFlowPerDayReq struct {
	SimCardIds *[]int64 `json:"sim_card_ids,omitempty"`

	Iccids *[]string `json:"iccids,omitempty"`

	// 月份
	Month *string `json:"month,omitempty"`

	// 日期
	Date *string `json:"date,omitempty"`
}

func (o SimCardFlowPerDayReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimCardFlowPerDayReq struct{}"
	}

	return strings.Join([]string{"SimCardFlowPerDayReq", string(data)}, " ")
}
