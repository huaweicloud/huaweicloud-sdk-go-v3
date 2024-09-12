package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimCardFlowItem struct {

	// sim卡标识
	SimCardId *int64 `json:"sim_card_id,omitempty"`

	// 容器ID
	Iccid *string `json:"iccid,omitempty"`

	// 流量
	Flow *float64 `json:"flow,omitempty"`
}

func (o SimCardFlowItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimCardFlowItem struct{}"
	}

	return strings.Join([]string{"SimCardFlowItem", string(data)}, " ")
}
