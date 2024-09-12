package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimCardFlowPerDayRsp struct {

	// 日期
	Date *string `json:"date,omitempty"`

	SimFlow *[]SimCardFlowItem `json:"sim_flow,omitempty"`
}

func (o SimCardFlowPerDayRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimCardFlowPerDayRsp struct{}"
	}

	return strings.Join([]string{"SimCardFlowPerDayRsp", string(data)}, " ")
}
