package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListFlowBySimCardsReq struct {

	// iccid列表（三网卡不支持），最大支持50，且iccid和sim_card_id列表二选一
	Iccids *[]string `json:"iccids,omitempty"`

	// sim_card_id列表，最大支持50，且iccid和sim_card_id列表二选一
	SimCardIds *[]int64 `json:"sim_card_ids,omitempty"`
}

func (o ListFlowBySimCardsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowBySimCardsReq struct{}"
	}

	return strings.Join([]string{"ListFlowBySimCardsReq", string(data)}, " ")
}
